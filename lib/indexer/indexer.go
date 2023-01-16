package indexer

import (
	"github.com/anchore/syft/syft/pkg"
)

type Indexer interface {
	Index(string) ([]pkg.Package, error)
}

type BaseIndexer struct {
	Indexer
}
