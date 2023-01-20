/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/valllabh/security-cli/lib/cataloger"
	"github.com/valllabh/security-cli/lib/ui"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan imageURL",
	Short: "scan scm repo",
	Run: func(cmd *cobra.Command, args []string) {
		Scan(args[0])
	},
	Args: cobra.ExactArgs(1),
}

func Scan(userInput string) {

	_, bom, err := cataloger.Catalog(userInput)

	if err != nil {
		fmt.Println(err)
	}

	packages := bom.Artifacts.PackageCatalog.Sorted()

	for _, p := range packages {
		
		ui.Step("Found " + p.Name + ":" + p.Version)
		// res, err := graph.Store(p)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// fmt.Println(res)
	}

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
