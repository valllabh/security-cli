/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/valllabh/security-cli/lib/indexer"
	"github.com/valllabh/security-cli/lib/indexer/cargo"
	"github.com/valllabh/security-cli/lib/indexer/golang"
	"github.com/valllabh/security-cli/lib/indexer/maven"
	"github.com/valllabh/security-cli/lib/indexer/npm"

	"github.com/anchore/syft/syft/pkg"
	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index packageManager artifact",
	Short: "Command is used to index artifacts",
	Long:  "Arguments #1 argument is for package manager; #2 artifact name",
	Run: func(cmd *cobra.Command, args []string) {

		var packages, err = invokeIndexer(args[0], args[1])

		fmt.Println(packages)

		if err != nil {
			fmt.Println(err)
		}

	},
	Args: cobra.ExactArgs(2),
}

func invokeIndexer(indexerName string, pkg string) ([]pkg.Package, error) {
	var indexerObject indexer.Indexer

	switch indexerName {
	case "cargo", "rust":
		indexerObject = &cargo.CargoIndexer{}

	case "npm", "node":
		indexerObject = &npm.NpmIndexer{}
	case "go":
		indexerObject = &golang.GoIndexer{}
	case "maven":
		indexerObject = &maven.MavenIndexer{}
	}

	if indexerObject != nil {
		return indexerObject.Index(pkg)
	}

	return nil, errors.New("INVALID PACKAGE MANAGER NAME [" + indexerName + "]")
}

func init() {

	rootCmd.AddCommand(indexCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// indexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// indexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
