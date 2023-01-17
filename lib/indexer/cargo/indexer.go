package cargo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"github.com/valllabh/security-cli/lib/indexer"

	"github.com/anchore/syft/syft/pkg"
)

type CargoIndexer struct {
	indexer.BaseIndexer
}

func (i *CargoIndexer) Index(artifact string) ([]pkg.Package, error) {

	// Set the base URL for the crates.io API
	baseURL := "https://crates.io/api/v1/crates"

	// Build the package info URL
	packageInfoURL := fmt.Sprintf("%s/%s", baseURL, url.QueryEscape(artifact))

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
	var data = &CargoPackageInfo{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	// Print the package information
	for _, version := range data.Versions {
		fmt.Println("Checksum:", version.Checksum)
	}

	return convert(data)
}

func convert(packageInfo *CargoPackageInfo) ([]pkg.Package, error) {
	// TODO: conversion logic here
	return nil, nil
}
