package giturl

import (
	"fmt"

	giturl "github.com/whilp/git-urls"

	"github.com/kubescape/go-git-url/apis/azureapi"
	"github.com/kubescape/go-git-url/apis/githubapi"
	"github.com/kubescape/go-git-url/apis/gitlabapi"
	azureparserv1 "github.com/kubescape/go-git-url/azureparser/v1"
	githubparserv1 "github.com/kubescape/go-git-url/githubparser/v1"
	gitlabparserv1 "github.com/kubescape/go-git-url/gitlabparser/v1"
)

// NewGitURL get instance of git parser
func NewGitURL(fullURL string) (IGitURL, error) {
	hostUrl, err := getHost(fullURL)
	if err != nil {
		return nil, err
	}

	if githubparserv1.IsHostGitHub(hostUrl) {
		return githubparserv1.NewGitHubParserWithURL(fullURL)
	}
	if azureparserv1.IsHostAzure(hostUrl) {
		return azureparserv1.NewAzureParserWithURL(fullURL)
	}
	if gitlabparserv1.IsHostGitLab(hostUrl) {
		return gitlabparserv1.NewGitLabParserWithURL(fullURL)
	}
	return nil, fmt.Errorf("repository host '%s' not supported", hostUrl)
}

// NewGitAPI get instance of git api
func NewGitAPI(fullURL string) (IGitAPI, error) {
	hostUrl, err := getHost(fullURL)
	if err != nil {
		return nil, err
	}

	switch hostUrl {
	case githubapi.DEFAULT_HOST, githubapi.RAW_HOST:
		return githubparserv1.NewGitHubParserWithURL(fullURL)
	case gitlabapi.DEFAULT_HOST:
		return gitlabparserv1.NewGitLabParserWithURL(fullURL)
	case azureapi.DEFAULT_HOST, azureapi.DEV_HOST:
		return azureparserv1.NewAzureParserWithURL(fullURL)
	default:
		return nil, fmt.Errorf("repository host '%s' not supported", hostUrl)
	}
}

func getHost(fullURL string) (string, error) {
	parsedURL, err := giturl.Parse(fullURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Host, nil
}
