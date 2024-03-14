package giturl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGitURL(t *testing.T) {
	tests := []struct {
		name     string
		fullURL  string
		provider string
		owner    string
		repo     string
		branch   string
		url      string
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:     "parse github",
			fullURL:  "https://github.com/kubescape/go-git-url",
			provider: "github",
			owner:    "kubescape",
			repo:     "go-git-url",
			url:      "https://github.com/kubescape/go-git-url",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse github www",
			fullURL:  "https://www.github.com/kubescape/go-git-url",
			provider: "github",
			owner:    "kubescape",
			repo:     "go-git-url",
			url:      "https://github.com/kubescape/go-git-url",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse github ssh",
			fullURL:  "git@github.com:kubescape/go-git-url.git",
			provider: "github",
			owner:    "kubescape",
			repo:     "go-git-url",
			url:      "https://github.com/kubescape/go-git-url",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse azure",
			fullURL:  "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public",
			provider: "azure",
			owner:    "dwertent",
			repo:     "ks-testing-public",
			url:      "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse azure ssh",
			fullURL:  "git@ssh.dev.azure.com:v3/dwertent/ks-testing-public/ks-testing-public",
			provider: "azure",
			owner:    "dwertent",
			repo:     "ks-testing-public",
			url:      "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse bitbucket https",
			fullURL:  "https://bitbucket.org/matthyx/ks-testing-public.git",
			provider: "bitbucket",
			owner:    "matthyx",
			repo:     "ks-testing-public",
			url:      "https://bitbucket.org/matthyx/ks-testing-public",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse bitbucket ssh",
			fullURL:  "git@bitbucket.org:matthyx/ks-testing-public.git",
			provider: "bitbucket",
			owner:    "matthyx",
			repo:     "ks-testing-public",
			url:      "https://bitbucket.org/matthyx/ks-testing-public",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse gitlab",
			fullURL:  "https://gitlab.com/kubescape/testing",
			provider: "gitlab",
			owner:    "kubescape",
			repo:     "testing",
			url:      "https://gitlab.com/kubescape/testing",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse gitlab ssh",
			fullURL:  "git@gitlab.com:kubescape/testing.git",
			provider: "gitlab",
			owner:    "kubescape",
			repo:     "testing",
			url:      "https://gitlab.com/kubescape/testing",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse gitlab branch",
			fullURL:  "https://gitlab.com/kubescape/testing/-/tree/dev",
			provider: "gitlab",
			owner:    "kubescape",
			repo:     "testing",
			branch:   "dev",
			url:      "https://gitlab.com/kubescape/testing",
			wantErr:  assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gh, err := NewGitURL(tt.fullURL)
			if !tt.wantErr(t, err, fmt.Sprintf("NewGitURL(%v)", tt.fullURL)) {
				return
			}
			assert.Equal(t, tt.provider, gh.GetProvider())
			assert.Equal(t, tt.owner, gh.GetOwnerName())
			assert.Equal(t, tt.repo, gh.GetRepoName())
			assert.Equal(t, tt.branch, gh.GetBranchName())
			assert.Equal(t, tt.url, gh.GetURL().String())
		})
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

	{
		gitURL, err = NewGitAPI("https://gitlab.host.com/kubescape/testing")
		assert.NoError(t, err)
	}
}
