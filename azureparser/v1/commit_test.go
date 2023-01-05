package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLatestCommit(t *testing.T) {
	{
		az, err := NewAzureParserWithURL(urlA)
		assert.NoError(t, err)
		latestCommit, err := az.GetLatestCommit()
		assert.NoError(t, err)
		assert.Equal(t, "4502b9b51ee3ac1ea649bacfa0f48ebdeab05f4a", latestCommit.SHA)
		assert.Equal(t, "GitHub", latestCommit.Committer.Name)
		assert.Equal(t, "", latestCommit.Committer.Email)
	}
}
