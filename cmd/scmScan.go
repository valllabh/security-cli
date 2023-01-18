/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/anchore/stereoscope/pkg/image"
	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/linux"
	"github.com/anchore/syft/syft/pkg/cataloger"

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

	local, _ := ioutil.TempDir("", "go-vcs")

	cfg := ProviderConfig{
		SyftProviderConfig: SyftProviderConfig{
			CatalogingOptions: cataloger.DefaultConfig(),
		},
	}
	_, sbom, err := syftProvider(remote, cfg)

	pacakges := sbom.Artifacts.PackageCatalog.Sorted()

	for _, p := range pacakges {
		fmt.Println(p.Name, p.Type)
	}

	bytes, err := syft.Encode(*sbom, syft.FormatByName("spdx-2-json"))
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create(local + "/spdx.json")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(string(bytes))

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(local + "/spdx.json" + " done")

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
