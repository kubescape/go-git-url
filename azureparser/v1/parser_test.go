package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	urlA = "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public"
	urlB = "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public?path=/rules-tests/alert-rw-hostpath/deployment/expected.json"
	urlC = "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public?path=/scripts&version=GBdev&_a=contents"
	urlD = "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public?path=/scripts&version=GTv1.0.1&_a=contents"
)

func TestNewGitHubParserWithURL(t *testing.T) {
	{
		az, err := NewAzureParserWithURL(urlA)
		assert.NoError(t, err)
		assert.Equal(t, "dev.azure.com", az.GetHostName())
		assert.Equal(t, "azure", az.GetProvider())
		assert.Equal(t, "dwertent", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, urlA, az.GetURL().String())
		assert.Equal(t, "", az.GetBranchName())
		assert.Equal(t, "", az.GetPath())
	}
	{
		az, err := NewAzureParserWithURL(urlB)
		assert.NoError(t, err)
		assert.Equal(t, "dev.azure.com", az.GetHostName())
		assert.Equal(t, "azure", az.GetProvider())
		assert.Equal(t, "dwertent", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, urlA, az.GetURL().String())
		assert.Equal(t, "", az.GetBranchName())
		assert.Equal(t, "/rules-tests/alert-rw-hostpath/deployment/expected.json", az.GetPath())
	}
	{
		az, err := NewAzureParserWithURL(urlC)
		assert.NoError(t, err)
		assert.Equal(t, "dev.azure.com", az.GetHostName())
		assert.Equal(t, "azure", az.GetProvider())
		assert.Equal(t, "dwertent", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, urlA, az.GetURL().String())
		assert.Equal(t, "dev", az.GetBranchName())
		assert.Equal(t, "", az.GetTag())
		assert.Equal(t, "/scripts", az.GetPath())
	}
	{
		az, err := NewAzureParserWithURL(urlD)
		assert.NoError(t, err)
		assert.Equal(t, "dev.azure.com", az.GetHostName())
		assert.Equal(t, "azure", az.GetProvider())
		assert.Equal(t, "dwertent", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, urlA, az.GetURL().String())
		assert.Equal(t, "v1.0.1", az.GetTag())
		assert.Equal(t, "", az.GetBranchName())
		assert.Equal(t, "/scripts", az.GetPath())
	}
}
