package gitlabapi

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubRoundTripper struct {
	body string
}

func (s *stubRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     "200 OK",
		Body:       io.NopCloser(strings.NewReader(s.body)),
		Header:     http.Header{},
		Request:    req,
	}, nil
}

func gitLabAPIReturning(body string) *GitLabAPI {
	gl := NewGitLabAPI("gitlab.com")
	gl.httpClient = &http.Client{Transport: &stubRoundTripper{body: body}}
	return gl
}

func TestGetLatestCommitEmptyList(t *testing.T) {
	// GitLab answers 200 with an empty array for a project that has no commits
	commit, err := gitLabAPIReturning(`[]`).GetLatestCommit("kubescape", "go-git-url", "master", &Headers{})

	assert.Error(t, err)
	assert.Nil(t, commit)
}

func TestGetLatestCommitNotAList(t *testing.T) {
	// anything that is not a JSON array leaves data nil
	commit, err := gitLabAPIReturning(`{"message":"401 Unauthorized"}`).GetLatestCommit("kubescape", "go-git-url", "master", &Headers{})

	assert.Error(t, err)
	assert.Nil(t, commit)
}

func TestGetLatestCommit(t *testing.T) {
	body := `[{"id":"e7d287e491b4002bc59d67ad7423d8119fc89e6c","short_id":"e7d287e4","title":"updated interface","author_name":"David Wertenteil"}]`

	commit, err := gitLabAPIReturning(body).GetLatestCommit("kubescape", "go-git-url", "master", &Headers{})

	assert.NoError(t, err)
	assert.NotNil(t, commit)
	assert.Equal(t, "e7d287e491b4002bc59d67ad7423d8119fc89e6c", commit.ID)
	assert.Equal(t, "updated interface", commit.Title)
}
