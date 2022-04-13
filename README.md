# GIT Parser

The `git-parser` is a package meant for parsing git urls

This package also enables listing all files based on there extension

> The package currently supports only `github` parser and API. Feel free to contribute any other
## Parser

### Parse a git URL

```go
package main

import (
	"fmt"

	giturl "github.com/armosec/go-git-url"
)

func main() {

	fullURl := "https://github.com/armosec/go-git-url"
	gitURL, err := giturl.NewGitURL(fullURl) // initialize and parse the URL
	if err != nil {
		// do something
	}

	fmt.Printf(gitURL.GetHost())  // github.com
	fmt.Printf(gitURL.GetOwner()) // armosec
	fmt.Printf(gitURL.GetRepo())  // url-git-go
}
```

## Git API support

> It is recommended to use a [github token](https://docs.github.com/en/enterprise-server@3.4/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token). Set the github token in the `GITHUB_TOKEN` env

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

### Download files
```go

	// Download all files
	all, err := gitURL.DownloadAllFiles()

	// Download all files with the listed extensions
	extensions := []string{"yaml", "json"}
	files, err := gitURL.DownloadFilesWithExtension(extensions)

```	 