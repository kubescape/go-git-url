package githubapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAll(t *testing.T) {
	tree, _ := NewMockGitHubAPI().GetRepoTree("armosec", "go-git-url", "", nil)
	assert.Equal(t, 25, len(tree.ListAll()))
}

func TestListAllFiles(t *testing.T) {
	tree, _ := NewMockGitHubAPI().GetRepoTree("armosec", "go-git-url", "", nil)
	assert.Equal(t, 19, len(tree.ListAllFiles()))
}

func TestListAllDirs(t *testing.T) {
	tree, _ := NewMockGitHubAPI().GetRepoTree("armosec", "go-git-url", "", nil)
	assert.Equal(t, 6, len(tree.ListAllDirs()))
}
