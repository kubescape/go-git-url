package azureapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAll(t *testing.T) {
	tree, _ := NewMockAzureAPI().GetRepoTree("anubhav06", "testing", "testing", "master", nil)
	assert.Equal(t, 12, len(tree.ListAll()))
}	

func TestListAllFiles(t *testing.T) {
	tree, _ := NewMockAzureAPI().GetRepoTree("anubhav06", "testing", "testing", "master", nil)
	assert.Equal(t, 11, len(tree.ListAllFiles()))
}

func TestListAllDirs(t *testing.T) {
	tree, _ := NewMockAzureAPI().GetRepoTree("anubhav06", "testing", "testing", "master", nil)
	assert.Equal(t, 1, len(tree.ListAllDirs()))
}
