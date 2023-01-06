package azureapi

import (
	"encoding/json"
	"fmt"
)

var mockTreeUrl = `{ "count": 12, "value": [ { "objectId": "3838adf7667bcec12a8bd0b868e5ce18e4476815", "gitObjectType": "tree", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/", "isFolder": true, "url": "https://dev.azure.com/anubhav06/2773a1e0-3a09-4b3c-962d-31b884fe58c4/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items?path=%2F&versionType=Branch&version=dev&versionOptions=None" }, { "objectId": "62c893550adb53d3a8fc29a1584ff831cb829062", "gitObjectType": "blob", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/.gitignore", "url": "https://dev.azure.com/anubhav06/testing/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items//.gitignore?versionType=Branch&version=dev&versionOptions=None" }, { "objectId": "9172c84f0c2141833cc3b0ed512f94e36e789aa0", "gitObjectType": "blob", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/README.md", "url": "https://dev.azure.com/anubhav06/testing/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items//README.md?versionType=Branch&version=dev&versionOptions=None" }, { "objectId": "ea6048a012f2a69a09660cea3138ac64db49c021", "gitObjectType": "blob", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/postgres-pod.yml", "url": "https://dev.azure.com/anubhav06/testing/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items//postgres-pod.yml?versionType=Branch&version=dev&versionOptions=None" }, { "objectId": "2f16bbce9eee40d071aed4bd9a0e601fc99b9216", "gitObjectType": "blob", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/postgres-service.yml", "url": "https://dev.azure.com/anubhav06/testing/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items//postgres-service.yml?versionType=Branch&version=dev&versionOptions=None" }, { "objectId": "a0af8b4a8fe692a4526a159e851a7bf7efde1c69", "gitObjectType": "blob", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/redis-pod.yml", "url": "https://dev.azure.com/anubhav06/testing/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items//redis-pod.yml?versionType=Branch&version=dev&versionOptions=None" }, { "objectId": "2bab8621dfb18f40b5487df40aaef3921e61849b", "gitObjectType": "blob", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/redis-service.yml", "url": "https://dev.azure.com/anubhav06/testing/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items//redis-service.yml?versionType=Branch&version=dev&versionOptions=None" }, { "objectId": "009faa3b128a0375772f4914b74419da3abd079b", "gitObjectType": "blob", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/result-app-pod.yml", "url": "https://dev.azure.com/anubhav06/testing/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items//result-app-pod.yml?versionType=Branch&version=dev&versionOptions=None" }, { "objectId": "523ad3a6ac0419d335ffa3269bc4e057d06fe63e", "gitObjectType": "blob", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/result-app-service.yml", "url": "https://dev.azure.com/anubhav06/testing/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items//result-app-service.yml?versionType=Branch&version=dev&versionOptions=None" }, { "objectId": "bcced3583f5c8a208ecf7149a195ad21d17083aa", "gitObjectType": "blob", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/voting-app-pod.yml", "url": "https://dev.azure.com/anubhav06/testing/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items//voting-app-pod.yml?versionType=Branch&version=dev&versionOptions=None" }, { "objectId": "d055a91baadb058e7ed211705f5e4b8f42127215", "gitObjectType": "blob", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/voting-app-service.yml", "url": "https://dev.azure.com/anubhav06/testing/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items//voting-app-service.yml?versionType=Branch&version=dev&versionOptions=None" }, { "objectId": "ea912ca1e7a8bf4bd4b149d68f97f692c3a11092", "gitObjectType": "blob", "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "path": "/worker-app-pod.yml", "url": "https://dev.azure.com/anubhav06/testing/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/items//worker-app-pod.yml?versionType=Branch&version=dev&versionOptions=None" } ] }`
var mockLastCommit = `{ "count": 1, "value": [ { "commitId": "8cb368a261285f7d5129109f9bb3918de1999449", "author": { "name": "Anubhav Gupta", "email": null, "date": "2022-12-22T12:58:10Z" }, "committer": { "name": "Anubhav Gupta", "email": null, "date": "2022-12-22T12:58:10Z" }, "comment": "Updated README.md", "changeCounts": { "Add": 0, "Edit": 1, "Delete": 0 }, "url": "https://dev.azure.com/anubhav06/2773a1e0-3a09-4b3c-962d-31b884fe58c4/_apis/git/repositories/e4a1a3d3-c2e7-4781-89af-2dff3a8412f6/commits/8cb368a261285f7d5129109f9bb3918de1999449", "remoteUrl": "https://dev.azure.com/anubhav06/testing/_git/testing/commit/8cb368a261285f7d5129109f9bb3918de1999449" } ] }`

type MockAzureAPI struct {
}

func NewMockAzureAPI() *MockAzureAPI { return &MockAzureAPI{} }

func (az *MockAzureAPI) GetRepoTree(owner, project, repo, branch string, headers *Headers) (*Tree, error) {
	t := Tree{}
	switch fmt.Sprintf("%s/%s/_git/%s", owner, project, repo) {
	case "anubhav06/testing/_git/testing":
		json.Unmarshal([]byte(mockTreeUrl), &t)
	}
	return &t, nil
}

func (az MockAzureAPI) GetDefaultBranchName(owner, project, repo string, headers *Headers) (string, error) {
	return "master", nil
}

func (az MockAzureAPI) GetLatestCommit(owner, project, repo, branch string, headers *Headers) (*Commit, error) {

	var data Commit

	if err := json.Unmarshal([]byte(mockLastCommit), &data); err != nil {
		return &data, err
	}
	return &data, nil
}
