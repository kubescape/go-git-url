package bitbucketapi

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

func bitBucketAPIReturning(body string) *BitBucketAPI {
	api := NewBitBucketAPI()
	api.httpClient = &http.Client{Transport: &stubRoundTripper{body: body}}
	return api
}

func TestGetLatestCommitNoValues(t *testing.T) {
	// a paged endpoint answers with an empty values array rather than an error
	commit, err := bitBucketAPIReturning(`{"values":[],"pagelen":10}`).GetLatestCommit("matthyx", "ks-testing-public", "main", &Headers{})

	assert.Error(t, err)
	assert.Nil(t, commit)
}

func TestGetLatestCommit(t *testing.T) {
	body := `{"values":[{"type":"commit","hash":"a1b2c3d4","message":"initial commit"}],"pagelen":10}`

	commit, err := bitBucketAPIReturning(body).GetLatestCommit("matthyx", "ks-testing-public", "main", &Headers{})

	assert.NoError(t, err)
	assert.NotNil(t, commit)
	assert.Equal(t, "a1b2c3d4", commit.Hash)
	assert.Equal(t, "initial commit", commit.Message)
}
