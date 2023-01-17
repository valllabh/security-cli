package maven

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"github.com/valllabh/security-cli/lib/indexer"

	"github.com/anchore/syft/syft/pkg"
)

type MavenIndexer struct {
	indexer.BaseIndexer
}

func (i *MavenIndexer) Index(artifact string) ([]pkg.Package, error) {

	// Set the base URL for the crates.io API
	baseURL := "https://deps.dev/_/s/maven/p"

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
	var data = &MavenPackageInfo{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	// Print the package information
	fmt.Println(data.Package.Name)

	return convert(data)
}

func convert(packageInfo *MavenPackageInfo) ([]pkg.Package, error) {
	// TODO: conversion logic here
	return nil, nil
}
