package apis

import "errors"

type ProviderType string

const (
	ProviderGitHub    ProviderType = "github"
	ProviderAzure     ProviderType = "azure"
	ProviderBitBucket ProviderType = "bitbucket"
	ProviderGitLab    ProviderType = "gitlab"
)

func (pt ProviderType) IsSupported() error {
	switch pt {
	case ProviderGitHub, ProviderAzure, ProviderBitBucket, ProviderGitLab:
		return nil
	}
	return errors.New("unsupported provider")
}

func (pt ProviderType) String() string {
	return string(pt)
}
