package apis

import "errors"

type ProviderType string

const (
	ProviderGitHub ProviderType = "github"
	ProviderAzure  ProviderType = "azure"
)

func (pt ProviderType) IsSupported() error {
	switch pt {
	case ProviderGitHub, ProviderAzure:
		return nil
	}
	return errors.New("unsupported provider")
}

func (pt ProviderType) String() string {
	return string(pt)
}
