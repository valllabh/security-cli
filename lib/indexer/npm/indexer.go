package npm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"valllabh/security-cli/lib/indexer"

	"github.com/anchore/syft/syft/pkg"
)

type NpmIndexer struct {
	indexer.BaseIndexer
}

func (i *NpmIndexer) Index(artifact string) ([]pkg.Package, error) {

	// Set the base URL for the crates.io API
	baseURL := "https://registry.npmjs.org"

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
	var data = &NpmPackageInfo{}

	err = json.Unmarshal(body, data)
	// err = json.NewDecoder(resp.Body).Decode(&packageInfo)

	if err != nil {
		return nil, err
	}

	// Print the package information
	for _, version := range data.Versions {
		fmt.Println("Checksum:", version.Version)
	}

	return convert(data)
}

func convert(packageInfo *NpmPackageInfo) ([]pkg.Package, error) {
	// TODO: conversion logic here
	return nil, nil
}
