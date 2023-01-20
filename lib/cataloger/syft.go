package cataloger

import (
	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/pkg/cataloger"
	"github.com/anchore/syft/syft/sbom"
	"github.com/anchore/syft/syft/source"
	"github.com/valllabh/security-cli/lib/ui"
)

func Catalog(userInput string) (Context, *sbom.SBOM, error) {

	ui.Step("Cataloging " + userInput)

	config := ProviderConfig{
		SyftProviderConfig: SyftProviderConfig{
			CatalogingOptions: cataloger.DefaultConfig(),
		},
	}

	if config.CatalogingOptions.Search.Scope == "" {
		return Context{}, nil, errDoesNotProvide
	}

	sourceInput, err := source.ParseInput(userInput, config.Platform, true)
	if err != nil {
		return Context{}, nil, err
	}

	src, cleanup, err := source.New(*sourceInput, config.RegistryOptions, config.Exclusions)
	if err != nil {
		return Context{}, nil, err
	}
	defer cleanup()

	catalog, relationships, theDistro, err := syft.CatalogPackages(src, config.CatalogingOptions)
	if err != nil {
		return Context{}, nil, err
	}
	context := Context{
		Source: &src.Metadata,
		Distro: theDistro,
	}

	sbom := &sbom.SBOM{
		Source:        src.Metadata,
		Relationships: relationships,
		Artifacts: sbom.Artifacts{
			PackageCatalog: catalog,
		},
	}

	ui.Step("Cataloging Done")

	return context, sbom, nil
}
