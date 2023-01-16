package golang

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"valllabh/security-cli/lib/indexer"

	"github.com/anchore/syft/syft/pkg"
)

type GoIndexer struct {
	indexer.BaseIndexer
}

func (i *GoIndexer) Index(artifact string) ([]pkg.Package, error) {

	//---

	// Set the base URL for the crates.io API
	baseURL := "https://deps.dev/_/s/go/p"

	// Build the package info URL
	packageInfoURL := fmt.Sprintf("%s/%s/v/", baseURL, url.QueryEscape(artifact))

	// Execute the search request
	resp, err := http.Get(packageInfoURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response into a struct
	var data = &GoPackageInfo{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	// Print the package information
	for _, d := range data.Version.Projects {
		fmt.Println("Checksum:", d.Name)
	}

	//---

	return convert(data)
}

func convert(packageInfo *GoPackageInfo) ([]pkg.Package, error) {
	// TODO: conversion logic here
	return nil, nil
}
