# GIT Parser

The `git-parser` is a package meant for parsing git urls

This package also enables listing all files based on there extension

## Parser

### Supported parsers

* GitHub
* GitLab
* Azure

### Parse a git URL

```go
package main

import (
	"fmt"

	giturl "github.com/kubescape/go-git-url"
)

func main() {

    fullURl := "https://github.com/kubescape/go-git-url"
	gitURL, err := giturl.NewGitURL(fullURl) // initialize and parse the URL
	if err != nil {
		// do something
	}

	fmt.Printf(gitURL.GetHostName())  // github.com
	fmt.Printf(gitURL.GetOwnerName()) // kubescape
	fmt.Printf(gitURL.GetRepoName())  // go-git-url
}
```
 
## Git API support

### Supported APIs

* GitHub
> It is recommended to use a [GitHub token](https://docs.github.com/en/enterprise-server@3.4/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token). Set the GitHub token in the `GITHUB_TOKEN` env

* GitLab
> It is recommended to use a [GitLab token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html#create-a-personal-access-token). Set the GitLab token in the `GITLAB_TOKEN` env

* Azure
> It is recommended to use a [Azure token](https://learn.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops&tabs=Windows#create-a-pat). Set the Azure token in the `AZURE_TOKEN` env

### List files and directories
```go

// List all files and directories names
all, err := gitURL.ListAllNames()

// List all files names
files, err := gitURL.ListFilesNames()

// List all directories names
dirs, err := gitURL.ListDirsNames()

// List files names with the listed extensions
extensions := []string{"yaml", "json"}
files, err := gitURL.ListFilesNamesWithExtension(extensions)

```		

Different URL support ->
```go
 
basicURL, err := giturl.NewGitURL("https://github.com/kubescape/go-git-url") 
 
nestedURL, err := giturl.NewGitURL("https://github.com/kubescape/go-git-url/tree/master/files")  

fileApiURL, err := giturl.NewGitURL("https://github.com/kubescape/go-git-url/blob/master/files/file0.json")  

fileRawURL, err := giturl.NewGitURL("https://raw.githubusercontent.com/kubescape/go-git-url/master/files/file0.json") 

```
### Download files
```go

// Download all files
all, err := gitURL.DownloadAllFiles()

// Download all files with the listed extensions
extensions := []string{"yaml", "json"}
files, err := gitURL.DownloadFilesWithExtension(extensions)

```	 
