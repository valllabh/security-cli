package graph

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"log"

	"github.com/dgraph-io/dgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/dgraph-io/dgo/protos/api"
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
	host := "blue-surf-800003.grpc.us-east-1.aws.cloud.dgraph.io:443"
	key := "ZTc1NTIxMDZiNzNhOGVhODBjNjkxYjY5Y2NkNmQ2ZWE="

	// u, err := url.Parse(endpoint)
	// if err != nil {
	// 	return nil, err
	// }

	// urlParts := strings.SplitN(u.Host, ".", 2)

	// host := urlParts[0] + ".grpc." + urlParts[1] + ":" + slashPort
	pool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	creds := credentials.NewClientTLSFromCert(pool, "")
	return grpc.Dial(
		host,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&authCreds{key}),
	)
}

func NewClient() *dgo.Dgraph {

	conn, err := getConnection()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return dgo.NewDgraphClient(api.NewDgraphClient(conn))
}

func Store(obj any) (*api.Assigned, error) {
	graphAPI := NewClient()

	txn := graphAPI.NewTxn()

	pb, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	mu := &api.Mutation{
		SetJson: pb,
	}

	res, err := txn.Mutate(context.Background(), mu)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}
