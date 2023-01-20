/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/anchore/stereoscope/pkg/image"
	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/linux"
	"github.com/anchore/syft/syft/pkg"
	"github.com/anchore/syft/syft/pkg/cataloger"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/valllabh/security-cli/lib/graph"

	"github.com/anchore/syft/syft/sbom"
	"github.com/anchore/syft/syft/source"
	"github.com/spf13/cobra"
)

type ProviderConfig struct {
	SyftProviderConfig
	SynthesisConfig
}

type Context struct {
	Source *source.Metadata
	Distro *linux.Release
}

type SyftProviderConfig struct {
	CatalogingOptions             cataloger.Config
	RegistryOptions               *image.RegistryOptions
	Platform                      string
	Exclusions                    []string
	AttestationPublicKey          string
	AttestationIgnoreVerification bool
}

type SynthesisConfig struct {
	GenerateMissingCPEs bool
}

var errDoesNotProvide = fmt.Errorf("cannot provide packages from the given source")

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan imageURL",
	Short: "scan scm repo",
	Run: func(cmd *cobra.Command, args []string) {
		Scan(args[0])
	},
	Args: cobra.ExactArgs(1),
}

func Scan(remote string) {

	cfg := ProviderConfig{
		SyftProviderConfig: SyftProviderConfig{
			CatalogingOptions: cataloger.DefaultConfig(),
		},
	}
	_, bom, err := syftProvider(remote, cfg)

	if err != nil {
		fmt.Println(err)
	}

	packages := bom.Artifacts.PackageCatalog.Sorted()

	for _, p := range packages {
		fmt.Println(p.Name, p.Type, p.Version, p.Type, p.Language)
	}

}

func storePackage(p *pkg.Package) {

	graphAPI := graph.NewClient()

	txn := graphAPI.NewTxn()

	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		SetJson: pb,
	}
	res, err := txn.Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}
}

func syftProvider(userInput string, config ProviderConfig) (Context, *sbom.SBOM, error) {
	if config.CatalogingOptions.Search.Scope == "" {
		return Context{}, nil, errDoesNotProvide
	}

	sourceInput, err := source.ParseInput(userInput, config.Platform, true)
	if err != nil {
		return Context{}, nil, err
	}

	src, cleanup, err := source.New(*sourceInput, config.RegistryOptions, config.Exclusions)
	if err != nil {
		return Context{}, nil, err
	}
	defer cleanup()

	catalog, relationships, theDistro, err := syft.CatalogPackages(src, config.CatalogingOptions)
	if err != nil {
		return Context{}, nil, err
	}
	context := Context{
		Source: &src.Metadata,
		Distro: theDistro,
	}

	sbom := &sbom.SBOM{
		Source:        src.Metadata,
		Relationships: relationships,
		Artifacts: sbom.Artifacts{
			PackageCatalog: catalog,
		},
	}

	return context, sbom, nil
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
