/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Options struct {
	None string
}

var options = &Options{}

// cmd represents the version command
var cmd = &cobra.Command{
	Use:   "version",
	Short: "show the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("No Versions Yet!")
	},
}

func init() {
	rootCmd.AddCommand(cmd)

	cmd.Flags().StringVarP(&options.None, "none", "n", "text", "Test Option")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
