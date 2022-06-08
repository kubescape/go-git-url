package v1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	urlA = "https://github.com/armosec/go-git-url"
	urlB = "https://github.com/armosec/go-git-url/blob/master/files/file0.json"
	urlC = "https://github.com/armosec/go-git-url/tree/master/files"
	urlD = "https://raw.githubusercontent.com/armosec/go-git-url/master/files/file0.json"
	urlE = "git@github.com:armosec/go-git-url.git"
)

func TestNewGitHubParserWithURL(t *testing.T) {
	{
		gh, err := NewGitHubParserWithURL(urlA)
		assert.NoError(t, err)
		assert.Equal(t, "github.com", gh.GetHostName())
		assert.Equal(t, "github", gh.GetProvider())
		assert.Equal(t, "armosec", gh.GetOwnerName())
		assert.Equal(t, "go-git-url", gh.GetRepoName())
		assert.Equal(t, urlA, gh.GetURL().String())
		assert.Equal(t, "", gh.GetBranchName())
		assert.Equal(t, "", gh.GetPath())
	}
	{
		gh, err := NewGitHubParserWithURL(urlB)
		assert.NoError(t, err)
		assert.Equal(t, "github.com", gh.GetHostName())
		assert.Equal(t, "armosec", gh.GetOwnerName())
		assert.Equal(t, "go-git-url", gh.GetRepoName())
		assert.Equal(t, "master", gh.GetBranchName())
		assert.Equal(t, "files/file0.json", gh.GetPath())
		assert.Equal(t, urlA, gh.GetURL().String())
	}
	{
		gh, err := NewGitHubParserWithURL(urlC)
		assert.NoError(t, err)
		assert.Equal(t, "github.com", gh.GetHostName())
		assert.Equal(t, "armosec", gh.GetOwnerName())
		assert.Equal(t, "go-git-url", gh.GetRepoName())
		assert.Equal(t, "master", gh.GetBranchName())
		assert.Equal(t, "files", gh.GetPath())
		assert.Equal(t, urlA, gh.GetURL().String())
	}
	{
		gh, err := NewGitHubParserWithURL(urlD)
		assert.NoError(t, err)
		assert.Equal(t, "github.com", gh.GetHostName())
		assert.Equal(t, "armosec", gh.GetOwnerName())
		assert.Equal(t, "go-git-url", gh.GetRepoName())
		assert.Equal(t, "master", gh.GetBranchName())
		assert.Equal(t, "files/file0.json", gh.GetPath())
		assert.Equal(t, urlA, gh.GetURL().String())
	}
	{
		gh, err := NewGitHubParserWithURL(urlE)
		assert.NoError(t, err)
		assert.Equal(t, "github.com", gh.GetHostName())
		assert.Equal(t, "armosec", gh.GetOwnerName())
		assert.Equal(t, "go-git-url", gh.GetRepoName())
		assert.Equal(t, urlA, gh.GetURL().String())
		assert.Equal(t, "", gh.GetBranchName())
		assert.Equal(t, "", gh.GetPath())
	}
}

func TestSetDefaultBranch(t *testing.T) {
	{
		gh, err := NewGitHubParserWithURL(urlA)
		assert.NoError(t, err)
		assert.NoError(t, gh.SetDefaultBranchName())
		assert.Equal(t, "master", gh.GetBranchName())
	}
	{
		gh, err := NewGitHubParserWithURL(strings.ReplaceAll(urlB, "master", "dev"))
		assert.NoError(t, err)
		assert.NoError(t, gh.SetDefaultBranchName())
		assert.Equal(t, "master", gh.GetBranchName())
	}
	{
		gh, err := NewGitHubParserWithURL(urlE)
		assert.NoError(t, err)
		assert.NoError(t, gh.SetDefaultBranchName())
		assert.Equal(t, "master", gh.GetBranchName())
	}
}
