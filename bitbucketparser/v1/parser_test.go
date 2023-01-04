package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	urlA = "https://bitbucket.org/matthyx/ks-testing-public"                                                        // general case
	urlB = "https://bitbucket.org/matthyx/ks-testing-public/src/master/rules/etcd-encryption-native/raw.rego"       // file
	urlC = "https://bitbucket.org/matthyx/ks-testing-public/src/dev/README.md"                                      // branch
	urlD = "https://bitbucket.org/matthyx/ks-testing-public/src/dev/"                                               // branch
	urlE = "https://bitbucket.org/matthyx/ks-testing-public/src/v1.0.178/README.md"                                 // TODO fix tag
	urlF = "https://bitbucket.org/matthyx/ks-testing-public/raw/4502b9b51ee3ac1ea649bacfa0f48ebdeab05f4a/README.md" // TODO fix sha
	// scp-like syntax supported by git for ssh
	// see: https://mirrors.edge.kernel.org/pub/software/scm/git/docs/git-clone.html#URLS
	// regular form
	urlG = "git@bitbucket.org:matthyx/ks-testing-public.git"
	// unexpected form: should not panic
	urlH = "git@bitbucket.org:matthyx/to/ks-testing-public.git"
)

func TestNewGitHubParserWithURL(t *testing.T) {
	{
		gl, err := NewBitBucketParserWithURL(urlA)
		assert.NoError(t, err)
		assert.Equal(t, "bitbucket.org", gl.GetHostName())
		assert.Equal(t, "bitbucket", gl.GetProvider())
		assert.Equal(t, "matthyx", gl.GetOwnerName())
		assert.Equal(t, "ks-testing-public", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "", gl.GetBranchName())
		assert.Equal(t, "", gl.GetPath())
	}
	{
		gl, err := NewBitBucketParserWithURL(urlB)
		assert.NoError(t, err)
		assert.Equal(t, "bitbucket.org", gl.GetHostName())
		assert.Equal(t, "bitbucket", gl.GetProvider())
		assert.Equal(t, "matthyx", gl.GetOwnerName())
		assert.Equal(t, "ks-testing-public", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "master", gl.GetBranchName())
		assert.Equal(t, "rules/etcd-encryption-native/raw.rego", gl.GetPath())
	}
	{
		gl, err := NewBitBucketParserWithURL(urlC)
		assert.NoError(t, err)
		assert.Equal(t, "bitbucket.org", gl.GetHostName())
		assert.Equal(t, "bitbucket", gl.GetProvider())
		assert.Equal(t, "matthyx", gl.GetOwnerName())
		assert.Equal(t, "ks-testing-public", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "dev", gl.GetBranchName())
		assert.Equal(t, "README.md", gl.GetPath())
	}
	{
		gl, err := NewBitBucketParserWithURL(urlD)
		assert.NoError(t, err)
		assert.Equal(t, "bitbucket.org", gl.GetHostName())
		assert.Equal(t, "bitbucket", gl.GetProvider())
		assert.Equal(t, "matthyx", gl.GetOwnerName())
		assert.Equal(t, "ks-testing-public", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "dev", gl.GetBranchName())
		assert.Equal(t, "", gl.GetPath())
	}
	{
		gl, err := NewBitBucketParserWithURL(urlE)
		assert.NoError(t, err)
		assert.Equal(t, "bitbucket.org", gl.GetHostName())
		assert.Equal(t, "bitbucket", gl.GetProvider())
		assert.Equal(t, "matthyx", gl.GetOwnerName())
		assert.Equal(t, "ks-testing-public", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "v1.0.178", gl.GetBranchName())
		assert.Equal(t, "README.md", gl.GetPath())
	}
	{
		gl, err := NewBitBucketParserWithURL(urlF)
		assert.NoError(t, err)
		assert.Equal(t, "bitbucket.org", gl.GetHostName())
		assert.Equal(t, "bitbucket", gl.GetProvider())
		assert.Equal(t, "matthyx", gl.GetOwnerName())
		assert.Equal(t, "ks-testing-public", gl.GetRepoName())
		assert.Equal(t, urlA, gl.GetURL().String())
		assert.Equal(t, "4502b9b51ee3ac1ea649bacfa0f48ebdeab05f4a", gl.GetBranchName())
		assert.Equal(t, "README.md", gl.GetPath())
	}
	{
		gl, err := NewBitBucketParserWithURL(urlG)
		assert.NoError(t, err)
		assert.Equal(t, "bitbucket.org", gl.GetHostName())
		assert.Equal(t, "bitbucket", gl.GetProvider())
		assert.Equal(t, "matthyx", gl.GetOwnerName())
		assert.Equal(t, "ks-testing-public", gl.GetRepoName())
		assert.Equal(t, "https://bitbucket.org/matthyx/ks-testing-public", gl.GetURL().String())
		assert.Equal(t, "", gl.GetBranchName())
		assert.Equal(t, "", gl.GetPath())
	}
	{
		gl, err := NewBitBucketParserWithURL(urlH)
		assert.NoError(t, err)
		assert.Equal(t, "bitbucket.org", gl.GetHostName())
		assert.Equal(t, "bitbucket", gl.GetProvider())
		assert.Equal(t, "matthyx", gl.GetOwnerName())
		assert.Equal(t, "to", gl.GetRepoName())
		assert.Equal(t, "https://bitbucket.org/matthyx/to", gl.GetURL().String())
		assert.Equal(t, "ks-testing-public.git", gl.GetBranchName()) // invalid input leads to incorrect guess. At least this does not panic.
		assert.Equal(t, "", gl.GetPath())
	}
}
