package v1

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kubescape/go-git-url/apis/gitlabapi"
	"k8s.io/utils/strings/slices"
)

// ListAll list all directories and files in url tree
func (gl *GitLabURL) GetTree() (*gitlabapi.Tree, error) {

	if gl.isFile {
		return &gitlabapi.Tree{
			gitlabapi.InnerTree{
				Path: gl.GetPath(),
				Type: gitlabapi.ObjectTypeFile,
			},
		}, nil
	}

	if gl.GetHostName() == "" || gl.GetOwnerName() == "" || gl.GetRepoName() == "" {
		return nil, fmt.Errorf("missing host/owner/repo")
	}

	if gl.GetBranchName() == "" {
		if err := gl.SetDefaultBranchName(); err != nil {
			return nil, fmt.Errorf("failed to get default branch. reason: %s", err.Error())
		}
	}
	
	repoTree, err := gl.gitLabAPI.GetRepoTree(gl.GetOwnerName(), gl.GetRepoName(), gl.GetBranchName(), gl.headers())
	if err != nil {
		return repoTree, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}

	return repoTree, nil
}

// ListAllNames list all directories and files in url tree
func (gl *GitLabURL) ListAllNames() ([]string, error) {
	repoTree, err := gl.GetTree()
	if err != nil {
		return []string{}, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}
	return repoTree.ListAll(), nil
}

// ListDirsNames list all directories in url tree
func (gl *GitLabURL) ListDirsNames() ([]string, error) {
	repoTree, err := gl.GetTree()
	if err != nil {
		return []string{}, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}

	return repoTree.ListAllDirs(), nil
}


// ListAll list all files in url tree
func (gl *GitLabURL) ListFilesNames() ([]string, error) {
	repoTree, err := gl.GetTree()
	if err != nil {
		return []string{}, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}

	return repoTree.ListAllFiles(), nil
}

// ListFilesNamesWithExtension list all files in path with the desired extension
func (gl *GitLabURL) ListFilesNamesWithExtension(filesExtensions []string) ([]string, error) {

	fileNames, err := gl.ListFilesNames()
	if err != nil {
		return []string{}, err
	}
	if len(fileNames) == 0 {
		return fileNames, nil
	}

	var files []string
	for _, path := range fileNames {
		if gl.GetPath() != "" && !strings.HasPrefix(path, gl.GetPath()) {
			continue
		}
		if slices.Contains(filesExtensions, getFileExtension(path)) {
			files = append(files, path)
		}

	}
	return files, nil
}

func getFileExtension(path string) string {
	return strings.TrimPrefix(filepath.Ext(path), ".")
}
