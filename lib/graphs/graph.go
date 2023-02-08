package graphs

import (
	"context"
	"crypto/md5"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/anchore/syft/syft/pkg"
	"github.com/dgraph-io/dgo/v2"
	"github.com/valllabh/security-cli/lib/ui"
	"github.com/valllabh/security-cli/lib/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/dgraph-io/dgo/v2/protos/api"
)

type Node interface {
	getType() string
	setUid(string)
	getUid() string
}

type PackageVersion struct {
	Node
	ID    string   `json:"PackageVersion.id,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

type License struct {
	Node
	Uid   string   `json:"uid,omitempty"`
	ID    string   `json:"License.id,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

type ProgrammingLanguage struct {
	Node
	Uid   string   `json:"uid,omitempty"`
	ID    string   `json:"ProgrammingLanguage.id,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

type PackageType struct {
	Node
	Uid   string   `json:"uid,omitempty"`
	ID    string   `json:"PackageType.id,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

type Package struct {
	Node
	Uid       string                 `json:"uid,omitempty"`
	Name      string                 `json:"Package.name,omitempty"`
	Versions  []*PackageVersion      `json:"Package.versions,omitempty"`
	Licenses  []*License             `json:"Package.licenses,omitempty"`
	Languages []*ProgrammingLanguage `json:"Package.languages,omitempty"`
	Type      *PackageType           `json:"Package.type,omitempty"`
	URL       *string                `json:"Package.url,omitempty"`
	DType     []string               `json:"dgraph.type,omitempty"`
}

type authCreds struct {
	token string
}

func (a *authCreds) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"Authorization": a.token}, nil
}

func (a *authCreds) RequireTransportSecurity() bool {
	return true
}

func getConnection() (*grpc.ClientConn, error) {

	ui.Step("DB: GRPC: Connecting")

	// host := "localhost:9080"
	host := "blue-surf-800003.grpc.us-east-1.aws.cloud.dgraph.io:443"
	key := "ZTc1NTIxMDZiNzNhOGVhODBjNjkxYjY5Y2NkNmQ2ZWE="

	pool, err := x509.SystemCertPool()
	if err != nil {
		ui.Fatal(err)
		return nil, err
	}

	creds := credentials.NewClientTLSFromCert(pool, "")

	ui.Step("DB: GRPC: Dialing")

	return grpc.Dial(
		host,
		// grpc.WithInsecure(),
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&authCreds{key}),
	)
}

func NewClient() *dgo.Dgraph {

	conn, err := getConnection()

	conn.Connect()

	if err != nil {
		ui.Fatal(err)
		return nil
	}

	ui.Step("DB: Creating Client")

	return dgo.NewDgraphClient(api.NewDgraphClient(conn))
}

func getQueryVarBlock(varName string, k string, v string) string {
	return varName + ` as var(func: eq(` + k + `, "` + v + `"))`
}

func getVarName(text string) string {
	hash := md5.Sum([]byte(text))
	return "A" + hex.EncodeToString(hash[:])
}

func getUid(queryVars map[string]string, node Node) string {

	packageVarName := getVarName(node.getType() + value)
	queryVars[packageVarName] = getQueryVarBlock(packageVarName, "Package.name", value)
	return `uid(` + packageVarName + `)`
}

func StorePackages(ps []pkg.Package) (*api.Response, error) {

	ctx := context.Background()

	ui.Step("DB: Storing")

	addPackageInputs := []*Package{}

	queryVars := map[string]string{}

	for _, p := range ps {

		packageObj := Package{
			Name:      p.Name,
			Versions:  []*PackageVersion{},
			Licenses:  []*License{},
			Languages: []*ProgrammingLanguage{},
			Type: &PackageType{
				ID: string(p.Type),
			},
			URL:   &p.PURL,
			DType: []string{"Package"},
		}

		packageObj.Uid = getUid(queryVars, "Package", "Package.name", packageObj.Name)
		// packageVarName := getVarName("Package" + packageObj.Name)
		// queryVars[packageVarName] = getQueryVarBlock(packageVarName, "Package.name", packageObj.Name)
		// packageObj.Uid = `uid(` + packageVarName + `)`

		packageTypeVarName := getVarName("PackageType" + packageObj.Type.ID)
		queryVars[packageTypeVarName] = getQueryVarBlock(packageTypeVarName, "PackageType.id", packageObj.Type.ID)
		packageObj.Type.Uid = `uid(` + packageTypeVarName + `)`

		versionObj := PackageVersion{
			ID:    p.Name + "-" + p.Version,
			DType: []string{"PackageVersion"},
		}
		versionVarName := getVarName("PackageVersion" + versionObj.ID)
		queryVars[versionVarName] = getQueryVarBlock(versionVarName, "PackageVersion.id", versionObj.ID)
		versionObj.Uid = `uid(` + versionVarName + `)`

		packageObj.Versions = append(packageObj.Versions, &versionObj)

		for _, l := range p.Licenses {

			licenseObj := License{
				ID:    l,
				DType: []string{"License"},
			}
			licenseVarName := getVarName("License" + licenseObj.ID)
			queryVars[licenseVarName] = getQueryVarBlock(licenseVarName, "License.id", l)
			licenseObj.Uid = `uid(` + licenseVarName + `)`

			packageObj.Licenses = append(packageObj.Licenses, &licenseObj)
		}

		if len(string(p.Language)) > 0 {
			programmingLanguageObj := ProgrammingLanguage{
				ID:    string(p.Language),
				DType: []string{"ProgrammingLanguage"},
			}

			programmingLanguageVarName := getVarName("ProgrammingLanguage" + programmingLanguageObj.ID)
			queryVars[programmingLanguageVarName] = getQueryVarBlock(programmingLanguageVarName, "ProgrammingLanguage.ID", programmingLanguageObj.ID)
			programmingLanguageObj.Uid = `uid(` + programmingLanguageVarName + `)`

			packageObj.Languages = append(packageObj.Languages, &programmingLanguageObj)
		}

		addPackageInputs = append(addPackageInputs, &packageObj)

	}

	_, queryVarsValues := util.GetKeysOrValues(queryVars)
	query := `
		query {
			` + strings.Join(queryVarsValues, "\n") + `
		}
	`

	graphAPI := NewClient()

	txn := graphAPI.NewTxn()

	ui.Step("DB: New Transaction")

	pb, err := json.Marshal(addPackageInputs)

	util.WriteToFile(string(pb), "./payload.json")

	if err != nil {
		ui.Fatal(err)
		return nil, err
	}

	mu := &api.Mutation{
		SetJson:   pb,
		CommitNow: true,
	}

	ui.Step(query)

	req := &api.Request{
		Query:     query,
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}

	ui.Step("DB: Mutating")

	// res, err := txn.Mutate(ctx, mu)

	res, err := txn.Do(ctx, req)

	if err != nil {
		ui.Fatal(err)
		return nil, err
	}

	ui.Step("DB: Mutation Done")

	return res, nil
}
