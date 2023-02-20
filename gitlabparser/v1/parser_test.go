package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	urlA = "https://gitlab.com/kubescape/testing"                                                     // general case
	urlB = "https://gitlab.com/kubescape/testing/-/blob/main/stable/acs-engine-autoscaler/Chart.yaml" // file
	urlC = "https://gitlab.com/kubescape/testing/-/blob/dev/README.md"                                // branch
	urlD = "https://gitlab.com/kubescape/testing/-/tree/dev"                                          // branch
	urlE = "https://gitlab.com/kubescape/testing/-/blob/v0.0.0/README.md"                             // tag
	urlF = "https://gitlab.com/kubescape/testing/-/raw/main/stable/acs-engine-autoscaler/Chart.yaml"
	// scp-like syntax supported by git for ssh
	// see: https://mirrors.edge.kernel.org/pub/software/scm/git/docs/git-clone.html#URLS
	// regular form
	urlG = "git@gitlab.com:owner/repo.git"
	// project and subproject
	urlH = "https://gitlab.com/matthyx1/subgroup1/project1.git"
	urlI = "https://gitlab.com/matthyx1/subgroup1/subsubgroup1/project1.git"
)

func TestNewGitHubParserWithURL(t *testing.T) {
	{
		gl, err := NewGitLabParserWithURL(urlA)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab.com", gl.GetHostName())
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "kubescape", gl.GetOwnerName())
		assert.Equal(t, "testing", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "", gl.GetBranchName())
		assert.Equal(t, "", gl.GetPath())
	}
	{
		gl, err := NewGitLabParserWithURL(urlB)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab.com", gl.GetHostName())
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "kubescape", gl.GetOwnerName())
		assert.Equal(t, "testing", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "main", gl.GetBranchName())
		assert.Equal(t, "stable/acs-engine-autoscaler/Chart.yaml", gl.GetPath())
	}
	{
		gl, err := NewGitLabParserWithURL(urlC)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab.com", gl.GetHostName())
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "kubescape", gl.GetOwnerName())
		assert.Equal(t, "testing", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "dev", gl.GetBranchName())
		assert.Equal(t, "README.md", gl.GetPath())
	}
	{
		gl, err := NewGitLabParserWithURL(urlD)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab.com", gl.GetHostName())
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "kubescape", gl.GetOwnerName())
		assert.Equal(t, "testing", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "dev", gl.GetBranchName())
		assert.Equal(t, "", gl.GetPath())
	}
	{
		gl, err := NewGitLabParserWithURL(urlE)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab.com", gl.GetHostName())
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "kubescape", gl.GetOwnerName())
		assert.Equal(t, "testing", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "v0.0.0", gl.GetBranchName())
		assert.Equal(t, "README.md", gl.GetPath())
	}
	{
		gl, err := NewGitLabParserWithURL(urlF)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab.com", gl.GetHostName())
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "kubescape", gl.GetOwnerName())
		assert.Equal(t, "testing", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "main", gl.GetBranchName())
		assert.Equal(t, "stable/acs-engine-autoscaler/Chart.yaml", gl.GetPath())
	}
	{
		gl, err := NewGitLabParserWithURL(urlG)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab.com", gl.GetHostName())
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "owner", gl.GetOwnerName())
		assert.Equal(t, "repo", gl.GetRepoName())
		assert.Equal(t, "https://gitlab.com/owner/repo", gl.GetURL().String())
		assert.Equal(t, "", gl.GetBranchName())
		assert.Equal(t, "", gl.GetPath())
	}
	{
		gl, err := NewGitLabParserWithURL(urlH)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab.com", gl.GetHostName())
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "matthyx1/subgroup1", gl.GetOwnerName())
		assert.Equal(t, "project1", gl.GetRepoName())
		assert.Equal(t, "https://gitlab.com/matthyx1/subgroup1/project1", gl.GetURL().String())
		assert.Equal(t, "", gl.GetBranchName()) // invalid input leads to incorrect guess. At least this does not panic.
		assert.Equal(t, "", gl.GetPath())
	}
	{
		gl, err := NewGitLabParserWithURL(urlI)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab.com", gl.GetHostName())
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "matthyx1/subgroup1/subsubgroup1", gl.GetOwnerName())
		assert.Equal(t, "project1", gl.GetRepoName())
		assert.Equal(t, "https://gitlab.com/matthyx1/subgroup1/subsubgroup1/project1", gl.GetURL().String())
		assert.Equal(t, "", gl.GetBranchName()) // invalid input leads to incorrect guess. At least this does not panic.
		assert.Equal(t, "", gl.GetPath())
	}
}
