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
	"github.com/anchore/syft/syft/pkg/cataloger"

	"github.com/anchore/syft/syft/sbom"
	"github.com/anchore/syft/syft/source"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan imageURL",
	Short: "scan scm repo",
	Run: func(cmd *cobra.Command, args []string) {
		Scan(args[0])
	},
	Args: cobra.ExactArgs(2),
}

func Scan(remote string) {

	local, _ := ioutil.TempDir("", "go-vcs")

	img := &image.Image{}

	src, err := source.NewFromImage(img, remote)

	if err != nil {
		fmt.Print(err)
	}

	config := cataloger.DefaultConfig()

	catalog, _, distro, err := syft.CatalogPackages(&src, config)

	if err != nil {
		fmt.Print(err)
	}

	s := sbom.SBOM{
		Artifacts: sbom.Artifacts{
			PackageCatalog:    catalog,
			LinuxDistribution: distro,
		},
		Source: src.Metadata,
	}

	bytes, err := syft.Encode(s, syft.FormatByName("spdx-2-json"))
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

func init() {
	scmCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
