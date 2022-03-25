# GIT Parser

The `git-parser` is a package meant for parsing git urls

This package also enables listing all files based on there extension

## Usage

```go
import (
	"fmt"
    "github.com/armosec/url-git-go"
)

func main(){

    fullURl := "https://github.com/armosec/url-git-go"
    gitURL, err := NewGitURL(fullURl) // initialize and parse the URL
    if err != nil{
        fmt.Print(err)
        return
    }

    fmt.Printf(gitURL.GetHost()) // github.com
    fmt.Printf(gitURL.GetOwner()) // armosec
    fmt.Printf(gitURL.GetRepo()) // url-git-go 
    
    { // list only json and yaml files
        files, err := gitURL.ListFiles([]string{"yaml", "json"})
        if err != nil{
            fmt.Print(err)
            return
        }
    
        fmt.Printf(len(files)) // 5
    }

    { // list all files
        files, err := gitURL.ListFiles([]string{})
        if err != nil{
            fmt.Print(err)
            return
        }
    
        fmt.Printf(len(files))
    }

    // get the branch name. In this case it will be the default branch since it was not specified in the URL
    fmt.Printf(gitURL.GetBranch()) // master
}
```