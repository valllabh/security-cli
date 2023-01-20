package cataloger

import (
	"github.com/anchore/stereoscope/pkg/image"
	"github.com/anchore/syft/syft/linux"
	"github.com/anchore/syft/syft/pkg/cataloger"
	"github.com/anchore/syft/syft/source"
)

type ProviderConfig struct {
	SyftProviderConfig
	SynthesisConfig
}

type Context struct {
	Source *source.Metadata
	Distro *linux.Release
}

type SyftProviderConfig struct {
	CatalogingOptions             cataloger.Config
	RegistryOptions               *image.RegistryOptions
	Platform                      string
	Exclusions                    []string
	AttestationPublicKey          string
	AttestationIgnoreVerification bool
}

type SynthesisConfig struct {
	GenerateMissingCPEs bool
}
