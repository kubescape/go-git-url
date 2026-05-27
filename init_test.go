package giturl

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	gitlabparserv1 "github.com/kubescape/go-git-url/gitlabparser/v1"
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
			name:     "parse github with port",
			fullURL:  "https://github.com:443/kubescape/go-git-url",
			provider: "github",
			owner:    "kubescape",
			repo:     "go-git-url",
			url:      "https://github.com/kubescape/go-git-url",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse github with non-default port",
			fullURL:  "https://github.com:8443/kubescape/go-git-url",
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
			name:     "parse github ssh with protocol",
			fullURL:  "ssh://git@github.com/kubescape/go-git-url.git",
			provider: "github",
			owner:    "kubescape",
			repo:     "go-git-url",
			url:      "https://github.com/kubescape/go-git-url",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse github ssh with protocol and port",
			fullURL:  "ssh://git@github.com:22/kubescape/go-git-url.git",
			provider: "github",
			owner:    "kubescape",
			repo:     "go-git-url",
			url:      "https://github.com/kubescape/go-git-url",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse github ssh with protocol and non-default port",
			fullURL:  "ssh://git@github.com:2222/kubescape/go-git-url.git",
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
			name:     "parse azure ssh with protocol",
			fullURL:  "ssh://git@ssh.dev.azure.com/v3/dwertent/ks-testing-public/ks-testing-public",
			provider: "azure",
			owner:    "dwertent",
			repo:     "ks-testing-public",
			url:      "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse azure ssh with protocol and port",
			fullURL:  "ssh://git@ssh.dev.azure.com:22/v3/dwertent/ks-testing-public/ks-testing-public",
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
			name:     "parse bitbucket https with port",
			fullURL:  "https://bitbucket.org:443/matthyx/ks-testing-public.git",
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
			name:     "parse bitbucket ssh with protocol",
			fullURL:  "ssh://git@bitbucket.org/matthyx/ks-testing-public.git",
			provider: "bitbucket",
			owner:    "matthyx",
			repo:     "ks-testing-public",
			url:      "https://bitbucket.org/matthyx/ks-testing-public",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse bitbucket ssh with protocol and port",
			fullURL:  "ssh://git@bitbucket.org:22/matthyx/ks-testing-public.git",
			provider: "bitbucket",
			owner:    "matthyx",
			repo:     "ks-testing-public",
			url:      "https://bitbucket.org/matthyx/ks-testing-public",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse bitbucket ssh with protocol and non-default port",
			fullURL:  "ssh://git@bitbucket.org:2222/matthyx/ks-testing-public.git",
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
			name:     "parse gitlab with port",
			fullURL:  "https://gitlab.com:443/kubescape/testing",
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
			name:     "parse gitlab ssh with protocol",
			fullURL:  "ssh://git@gitlab.com/kubescape/testing.git",
			provider: "gitlab",
			owner:    "kubescape",
			repo:     "testing",
			url:      "https://gitlab.com/kubescape/testing",
			wantErr:  assert.NoError,
		},
		{
			name:     "parse gitlab ssh with protocol and port",
			fullURL:  "ssh://git@gitlab.com:22/kubescape/testing.git",
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
		{
			name:     "repo only",
			fullURL:  "https://git.host.com/repo.git",
			provider: "gitlab",
			owner:    "",
			repo:     "repo",
			url:      "https://git.host.com/repo",
			wantErr:  assert.NoError,
		},
		{
			name:     "repo only with port",
			fullURL:  "https://git.host.com:443/repo.git",
			provider: "gitlab",
			owner:    "",
			repo:     "repo",
			url:      "https://git.host.com/repo",
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

func TestNewGitURL_GitLabSelfHostedCustomPort(t *testing.T) {
	gitURL, err := NewGitURL("https://gitlab.host.com:8443/kubescape/testing")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "gitlab", gitURL.GetProvider())
	assert.Equal(t, "gitlab.host.com:8443", gitURL.GetHostName())
	assert.Equal(t, "kubescape", gitURL.GetOwnerName())
	assert.Equal(t, "testing", gitURL.GetRepoName())
	assert.Equal(t, "https://gitlab.host.com:8443/kubescape/testing", gitURL.GetURL().String())
	assert.Equal(t, "https://gitlab.host.com:8443/kubescape/testing.git", gitURL.GetHttpCloneURL())
}

func TestNewGitLabParserWithURL_IPv6HostPreservesBrackets(t *testing.T) {
	gitURL, err := gitlabparserv1.NewGitLabParserWithURL("", "https://[2001:db8::1]/kubescape/testing")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "[2001:db8::1]", gitURL.GetHostName())
	assert.Equal(t, "https://[2001:db8::1]/kubescape/testing", gitURL.GetURL().String())
	assert.Equal(t, "https://[2001:db8::1]/kubescape/testing.git", gitURL.GetHttpCloneURL())
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

func TestNewGitAPI_AzureCustomPortAPIRequestsUseParsedPort(t *testing.T) {
	requestTarget := make(chan string, 1)
	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		select {
		case requestTarget <- r.URL.Host:
		default:
		}
		return &http.Response{
			StatusCode: http.StatusBadGateway,
			Status:     "502 Bad Gateway",
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("blocked by test transport")),
			Request:    r,
		}, nil
	})
	defer func() {
		http.DefaultTransport = originalTransport
	}()

	gitURL, err := NewGitAPI("https://dev.azure.com:8443/dwertent/ks-testing-public/_git/ks-testing-public")
	if !assert.NoError(t, err) {
		return
	}

	assert.Error(t, gitURL.SetDefaultBranchName())

	select {
	case target := <-requestTarget:
		assert.Equal(t, "dev.azure.com:8443", target)
	default:
		t.Fatal("expected Azure API request to use the test transport")
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
