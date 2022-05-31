package v1

import (
	"os"
	"testing"

	"github.com/armosec/go-git-url/apis/githubapi"
	"github.com/stretchr/testify/assert"
)

func NewGitHubParserWithURLMock(fullURL string) (*GitHubURL, error) {
	gh := NewGitHubParserMock()

	if err := gh.Parse(fullURL); err != nil {
		return gh, err
	}

	return gh, nil
}

func NewGitHubParserMock() *GitHubURL {

	return &GitHubURL{
		gitHubAPI: githubapi.NewMockGitHubAPI(),
		host:      githubapi.DEFAULT_HOST,
		token:     os.Getenv("GITHUB_TOKEN"),
	}
}
func TestListAllNames(t *testing.T) {
	{
		gh, err := NewGitHubParserWithURLMock(urlA)
		assert.NoError(t, err)
		tree, err := gh.ListAllNames()
		assert.NoError(t, err)
		assert.Equal(t, 25, len(tree))
	}
}

func TestGetFileExtension(t *testing.T) {
	assert.Equal(t, "yaml", getFileExtension("my/name.yaml"))
	assert.Equal(t, "txt", getFileExtension("/my/name.txt"))
	assert.Equal(t, "json", getFileExtension("myName.json"))
	assert.Equal(t, "", getFileExtension("myName"))
}
