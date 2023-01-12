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
	urlE = "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public?path=%2F&version=GBdev"
	urlF = "https://dwertent@dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public"
	// scp-like syntax supported by git for ssh
	// see: https://mirrors.edge.kernel.org/pub/software/scm/git/docs/git-clone.html#URLS
	// regular form
	urlG = "git@ssh.dev.azure.com:v3/dwertent/ks-testing-public/ks-testing-public"
)

func TestNewAzureParserWithURL(t *testing.T) {
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
	{
		az, err := NewAzureParserWithURL(urlE)
		assert.NoError(t, err)
		assert.Equal(t, "dev.azure.com", az.GetHostName())
		assert.Equal(t, "azure", az.GetProvider())
		assert.Equal(t, "dwertent", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, urlA, az.GetURL().String())
		assert.Equal(t, "", az.GetTag())
		assert.Equal(t, "dev", az.GetBranchName())
		assert.Equal(t, "/", az.GetPath())
	}
	{
		az, err := NewAzureParserWithURL(urlF)
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
		az, err := NewAzureParserWithURL(urlG)
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, "dev.azure.com", az.GetHostName())
		assert.Equal(t, "azure", az.GetProvider())
		assert.Equal(t, "dwertent", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, urlA, az.GetURL().String())
		assert.Equal(t, "", az.GetBranchName())
		assert.Equal(t, "", az.GetPath())
	}
}

func TestSetDefaultBranch(t *testing.T) {
	{
		az, err := NewAzureParserWithURL(urlA)
		assert.NoError(t, err)
		assert.NoError(t, az.SetDefaultBranchName())
		assert.Equal(t, "master", az.GetBranchName())
	}
	{
		az, err := NewAzureParserWithURL(urlE)
		assert.NoError(t, err)
		assert.NoError(t, az.SetDefaultBranchName())
		assert.Equal(t, "master", az.GetBranchName())
	}
	{
		az, err := NewAzureParserWithURL(urlF)
		assert.NoError(t, err)
		assert.NoError(t, az.SetDefaultBranchName())
		assert.Equal(t, "master", az.GetBranchName())
	}
}
