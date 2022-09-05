package v1

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kubescape/go-git-url/apis/githubapi"
	"k8s.io/utils/strings/slices"
)

// ListAll list all directories and files in url tree
func (gh *GitHubURL) GetTree() (*githubapi.Tree, error) {
	if gh.isFile {
		return &githubapi.Tree{
			InnerTrees: []githubapi.InnerTree{
				{
					Path: gh.GetPath(),
					Type: githubapi.ObjectTypeFile,
				},
			},
		}, nil
	}
	if gh.GetHostName() == "" || gh.GetOwnerName() == "" || gh.GetRepoName() == "" {
		return nil, fmt.Errorf("missing host/owner/repo")
	}
	if gh.GetBranchName() == "" {
		if err := gh.SetDefaultBranchName(); err != nil {
			return nil, fmt.Errorf("failed to get default branch. reason: %s", err.Error())
		}
	}

	repoTree, err := gh.gitHubAPI.GetRepoTree(gh.GetOwnerName(), gh.GetRepoName(), gh.GetBranchName(), gh.headers())
	if err != nil {
		return repoTree, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}

	return repoTree, nil
}

// ListAllNames list all directories and files in url tree
func (gh *GitHubURL) ListAllNames() ([]string, error) {
	repoTree, err := gh.GetTree()
	if err != nil {
		return []string{}, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}
	return repoTree.ListAll(), nil
}

// ListDirsNames list all directories in url tree
func (gh *GitHubURL) ListDirsNames() ([]string, error) {
	repoTree, err := gh.GetTree()
	if err != nil {
		return []string{}, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}

	return repoTree.ListAllDirs(), nil
}

// ListAll list all files in url tree
func (gh *GitHubURL) ListFilesNames() ([]string, error) {
	repoTree, err := gh.GetTree()
	if err != nil {
		return []string{}, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}

	return repoTree.ListAllFiles(), nil
}

// ListFilesNamesWithExtension list all files in path with the desired extension
func (gh *GitHubURL) ListFilesNamesWithExtension(filesExtensions []string) ([]string, error) {

	fileNames, err := gh.ListFilesNames()
	if err != nil {
		return []string{}, err
	}
	if len(fileNames) == 0 {
		return fileNames, nil
	}

	var files []string
	for _, path := range fileNames {
		if gh.GetPath() != "" && !strings.HasPrefix(path, gh.GetPath()) {
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
