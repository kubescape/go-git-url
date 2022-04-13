package main

import (
	"fmt"
	"net/url"

	"github.com/armosec/go-git-url/apis/githubapi"
	githubparserv1 "github.com/armosec/go-git-url/githubparser/v1"
)

// NewGitURL get instance of git parser
func NewGitURL(fullURL string) (IGitURL, error) {
	hostUrl, err := getHost(fullURL)
	if err != nil {
		return nil, err
	}

	switch hostUrl {
	case githubapi.DEFAULT_HOST, githubapi.RAW_HOST:
		return githubparserv1.NewGitHubParserWithURL(fullURL)
	default:
		return nil, fmt.Errorf("repository host '%s' not supported", hostUrl)
	}
}

func getHost(fullURL string) (string, error) {
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Host, nil
}
