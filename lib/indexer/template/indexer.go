package changehere

import (
	"github.com/valllabh/security-cli/lib/indexer"

	"github.com/anchore/syft/syft/pkg"
)

type ChangeHereIndexer struct {
	indexer.BaseIndexer
}

func (i *ChangeHereIndexer) Index(artifact string) ([]pkg.Package, error) {

	var data = &ChangeHerePackageInfo{}

	return convert(data)
}

func convert(packageInfo *ChangeHerePackageInfo) ([]pkg.Package, error) {
	// TODO: conversion logic here
	return nil, nil
}
