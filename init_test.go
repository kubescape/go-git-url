package giturl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGitURL(t *testing.T) {
	{ // parse github
		const githubURL = "https://github.com/kubescape/go-git-url"
		gh, err := NewGitURL(githubURL)
		assert.NoError(t, err)
		assert.Equal(t, "github", gh.GetProvider())
		assert.Equal(t, "kubescape", gh.GetOwnerName())
		assert.Equal(t, "go-git-url", gh.GetRepoName())
		assert.Equal(t, "", gh.GetBranchName())
		assert.Equal(t, githubURL, gh.GetURL().String())
	}
	{ // parse github
		const githubURL = "https://www.github.com/kubescape/go-git-url"
		gh, err := NewGitURL(githubURL)
		assert.NoError(t, err)
		assert.Equal(t, "github", gh.GetProvider())
		assert.Equal(t, "kubescape", gh.GetOwnerName())
		assert.Equal(t, "go-git-url", gh.GetRepoName())
		assert.Equal(t, "", gh.GetBranchName())
		assert.Equal(t, "https://github.com/kubescape/go-git-url", gh.GetURL().String())
	}
	{ // parse github
		const githubURL = "git@github.com:kubescape/go-git-url.git"
		gh, err := NewGitURL(githubURL)
		assert.NoError(t, err)
		assert.Equal(t, "github", gh.GetProvider())
		assert.Equal(t, "kubescape", gh.GetOwnerName())
		assert.Equal(t, "go-git-url", gh.GetRepoName())
		assert.Equal(t, "", gh.GetBranchName())
		assert.Equal(t, "https://github.com/kubescape/go-git-url", gh.GetURL().String())
	}
	{ // parse azure
		const azureURL = "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public"
		az, err := NewGitURL(azureURL)
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, "azure", az.GetProvider())
		assert.Equal(t, "dwertent", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, "", az.GetBranchName())
		assert.Equal(t, "", az.GetPath())
		assert.Equal(t, azureURL, az.GetURL().String())
	}
	{ // parse azure
		const azureURL = "git@ssh.dev.azure.com:v3/dwertent/ks-testing-public/ks-testing-public"
		az, err := NewGitURL(azureURL)
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, "azure", az.GetProvider())
		assert.Equal(t, "dwertent", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, "", az.GetBranchName())
		assert.Equal(t, "", az.GetPath())
		assert.Equal(t, "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public", az.GetURL().String())
	}
	{ // parse bitbucket https
		az, err := NewGitURL("https://matthyx@bitbucket.org/matthyx/ks-testing-public.git")
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, "bitbucket", az.GetProvider())
		assert.Equal(t, "matthyx", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, "", az.GetBranchName())
		assert.Equal(t, "", az.GetPath())
		assert.Equal(t, "https://bitbucket.org/matthyx/ks-testing-public", az.GetURL().String())
	}
	{ // parse bitbucket ssh
		az, err := NewGitURL("git@bitbucket.org:matthyx/ks-testing-public.git")
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, "bitbucket", az.GetProvider())
		assert.Equal(t, "matthyx", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, "", az.GetBranchName())
		assert.Equal(t, "", az.GetPath())
		assert.Equal(t, "https://bitbucket.org/matthyx/ks-testing-public", az.GetURL().String())
	}
	{ // parse gitlab
		const gitlabURL = "https://gitlab.com/kubescape/testing"
		gl, err := NewGitURL(gitlabURL)
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "kubescape", gl.GetOwnerName())
		assert.Equal(t, "testing", gl.GetRepoName())
		assert.Equal(t, "", gl.GetBranchName())
		assert.Equal(t, "", gl.GetPath())
		assert.Equal(t, gitlabURL, gl.GetURL().String())
	}
	{ // parse gitlab
		const gitlabURL = "git@gitlab.com:kubescape/testing.git"
		gl, err := NewGitURL(gitlabURL)
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "kubescape", gl.GetOwnerName())
		assert.Equal(t, "testing", gl.GetRepoName())
		assert.Equal(t, "", gl.GetBranchName())
		assert.Equal(t, "", gl.GetPath())
		assert.Equal(t, "https://gitlab.com/kubescape/testing", gl.GetURL().String())
	}
	{ // parse gitlab
		const gitlabURL = "https://gitlab.com/kubescape/testing/-/tree/dev"
		gl, err := NewGitURL(gitlabURL)
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, "gitlab", gl.GetProvider())
		assert.Equal(t, "kubescape", gl.GetOwnerName())
		assert.Equal(t, "testing", gl.GetRepoName())
		assert.Equal(t, "dev", gl.GetBranchName())
		assert.Equal(t, "", gl.GetPath())
		assert.Equal(t, "https://gitlab.com/kubescape/testing", gl.GetURL().String())
	}

}
func TestNewGitAPI(t *testing.T) {
	fileText := "https://raw.githubusercontent.com/kubescape/go-git-url/master/files/file0.text"
	var gitURL IGitAPI
	var err error
	{
		gitURL, err = NewGitAPI("https://github.com/kubescape/go-git-url")
		assert.NoError(t, err)

		files, err := gitURL.ListFilesNamesWithExtension([]string{"yaml", "json"})
		assert.NoError(t, err)
		assert.Equal(t, 8, len(files))
	}

	{
		gitURL, err = NewGitAPI("https://github.com/kubescape/go-git-url")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))

	}

	{
		gitURL, err = NewGitAPI(fileText)
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))
	}

	{
		gitURL, err = NewGitAPI(fileText)
		assert.NoError(t, err)

		files, errM := gitURL.DownloadAllFiles()
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))
	}

	{
		gitURL, err := NewGitAPI("https://github.com/kubescape/go-git-url/tree/master/files")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))

	}

	{
		gitURL, err = NewGitAPI("https://github.com/kubescape/go-git-url/blob/master/files/file0.text")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))

	}
}
