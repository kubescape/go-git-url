package githubapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAll(t *testing.T) {
	tree, _ := NewMockGitHubAPI().GetRepoTree("armosec", "url-git-go", "", nil)
	assert.Equal(t, 25, tree.ListAll())
}

func TestListAllFiles(t *testing.T) {
	tree, _ := NewMockGitHubAPI().GetRepoTree("armosec", "url-git-go", "", nil)
	assert.Equal(t, 19, len(tree.ListAllFiles()))
}

func TestListAllDirs(t *testing.T) {
	tree, _ := NewMockGitHubAPI().GetRepoTree("armosec", "url-git-go", "", nil)
	assert.Equal(t, 6, len(tree.ListAllDirs()))
}
