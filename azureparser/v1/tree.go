package v1

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kubescape/go-git-url/apis/azureapi"
	"k8s.io/utils/strings/slices"
)

// ListAll list all directories and files in url tree
func (az *AzureURL) GetTree() (*azureapi.Tree, error) {

	if az.isFile {
		return &azureapi.Tree{
			InnerTree: []azureapi.InnerTree{
				{
					Path:          az.GetPath(),
					GitObjectType: azureapi.ObjectTypeFile,
				},
			},
		}, nil
	}

	if az.GetHostName() == "" || az.GetOwnerName() == "" || az.GetProjectName() == "" || az.GetRepoName() == "" {
		return nil, fmt.Errorf("missing host/owner/project/repo")
	}

	if az.GetBranchName() == "" {
		if err := az.SetDefaultBranchName(); err != nil {
			return nil, fmt.Errorf("failed to get default branch. reason: %s", err.Error())
		}
	}

	repoTree, err := az.azureAPI.GetRepoTree(az.GetOwnerName(), az.GetProjectName(), az.GetRepoName(), az.GetBranchName(), az.headers())
	if err != nil {
		return repoTree, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}

	return repoTree, nil
}

// ListAllNames list all directories and files in url tree
func (az *AzureURL) ListAllNames() ([]string, error) {

	repoTree, err := az.GetTree()

	if err != nil {
		return []string{}, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}

	return repoTree.ListAll(), nil
}

// ListDirsNames list all directories in url tree
func (az *AzureURL) ListDirsNames() ([]string, error) {
	repoTree, err := az.GetTree()
	if err != nil {
		return []string{}, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}

	return repoTree.ListAllDirs(), nil
}

// ListAll list all files in url tree
func (az *AzureURL) ListFilesNames() ([]string, error) {
	repoTree, err := az.GetTree()
	if err != nil {
		return []string{}, fmt.Errorf("failed to get repo tree. reason: %s", err.Error())
	}

	return repoTree.ListAllFiles(), nil
}

// ListFilesNamesWithExtension list all files in path with the desired extension
func (az *AzureURL) ListFilesNamesWithExtension(filesExtensions []string) ([]string, error) {

	fileNames, err := az.ListFilesNames()
	if err != nil {
		return []string{}, err
	}
	if len(fileNames) == 0 {
		return fileNames, nil
	}

	var files []string
	for _, path := range fileNames {
		if az.GetPath() != "" && !strings.HasPrefix(path, az.GetPath()) {
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
