package graphs

import (
	"context"
	"crypto/x509"
	"encoding/json"

	"github.com/anchore/syft/syft/pkg"
	"github.com/dgraph-io/dgo/v2"
	"github.com/valllabh/security-cli/graph/model"
	"github.com/valllabh/security-cli/lib/ui"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/dgraph-io/dgo/v2/protos/api"
)

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

func StorePackages(ps []pkg.Package) (*api.Response, error) {

	ctx := context.Background()

	ui.Step("DB: Storing")

	addPackageInputs := []*model.AddPackageInput{}

	for _, p := range ps {

		pr := model.AddPackageInput{
			Name:      p.Name,
			Versions:  []*model.PackageVersionRef{},
			Licenses:  []*model.LicenseRef{},
			Languages: []*model.ProgrammingLanguageRef{},
			Type: &model.PackageTypeRef{
				ID: string(p.Type),
			},
			URL: &p.PURL,
		}

		pr.Versions = append(pr.Versions, &model.PackageVersionRef{
			ID: p.Version,
		})

		for _, l := range p.Licenses {
			pr.Licenses = append(pr.Licenses, &model.LicenseRef{
				ID: l,
			})
		}

		pr.Languages = append(pr.Languages, &model.ProgrammingLanguageRef{
			ID: string(p.Language),
		})

		addPackageInputs = append(addPackageInputs, &pr)

	}

	graphAPI := NewClient()

	txn := graphAPI.NewTxn()

	ui.Step("DB: New Transaction")

	pb, err := json.Marshal(addPackageInputs)

	if err != nil {
		ui.Fatal(err)
		return nil, err
	}

	mu := &api.Mutation{
		SetJson:   pb,
		CommitNow: true,
	}

	ui.Step("DB: Mutating")

	res, err := txn.Mutate(ctx, mu)

	if err != nil {
		ui.Fatal(err)
		return nil, err
	}

	ui.Step("DB: Mutation Done")

	return res, nil
}
