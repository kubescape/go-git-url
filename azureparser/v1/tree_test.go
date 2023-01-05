package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAllNames(t *testing.T) {
	{
		az, err := NewAzureParserWithURL(urlA)
		assert.NoError(t, err)
		tree, err := az.ListAllNames()
		assert.NoError(t, err)
		assert.Equal(t, 1440, len(tree))
	}
}

func TestGetFileExtension(t *testing.T) {
	assert.Equal(t, "yaml", getFileExtension("my/name.yaml"))
	assert.Equal(t, "txt", getFileExtension("/my/name.txt"))
	assert.Equal(t, "json", getFileExtension("myName.json"))
	assert.Equal(t, "", getFileExtension("myName"))
}
