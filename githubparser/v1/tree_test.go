package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestListAllNames(t *testing.T) {
// 	{
// 		gh, err := NewGitHubParserWithURL(urlA)
// 		assert.NoError(t, err)
// 		tree, err := gh.ListAllNames()
// 		assert.NoError(t, err)
// 		assert.Less(t, 0, len(tree))
// 	}
// }

func TestGetFileExtension(t *testing.T) {
	assert.Equal(t, "yaml", getFileExtension("my/name.yaml"))
	assert.Equal(t, "txt", getFileExtension("/my/name.txt"))
	assert.Equal(t, "json", getFileExtension("myName.json"))
	assert.Equal(t, "", getFileExtension("myName"))
}
