package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLatestCommit(t *testing.T) {
	{
		gh, err := NewGitHubParserWithURLMock(urlA)
		assert.NoError(t, err)
		latestCommit, err := gh.GetLatestCommit()
		assert.NoError(t, err)
		assert.Equal(t, "e7d287e491b4002bc59d67ad7423d8119fc89e6c", latestCommit.SHA)
		assert.Equal(t, "David Wertenteil", latestCommit.Committer.Name)
		assert.Equal(t, "dwertent@armosec.io", latestCommit.Committer.Email)
	}
}
