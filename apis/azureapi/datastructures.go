package azureapi

import (
	"time"
)

type ObjectType string

const (
	ObjectTypeDir  ObjectType = "tree"
	ObjectTypeFile ObjectType = "blob"
)

type InnerTree struct {
	ObjectID      string     `json:"objectId"`
	GitObjectType ObjectType `json:"gitObjectType"`
	CommitID      string     `json:"commitId"`
	Path          string     `json:"path"`
	IsFolder      bool       `json:"isFolder"`
	Url           string     `json:"url"`
}

type Tree struct {
	Count     int         `json:"count"`
	InnerTree []InnerTree `json:"value"`
}

type CommitsMetadata struct {
	TreeID   string `json:"treeId"`
	CommitID string `json:"commitId"`
	Author   struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	} `json:"author"`
	Committer struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	} `json:"committer"`
	Comment string   `json:"comment"`
	Parents []string `json:"parents"`
	URL     string   `json:"url"`
}

type BranchValue struct {
	Commit        CommitsMetadata `json:"commit"`
	Name          string          `json:"name"`
	AheadCount    int             `json:"aheadCount"`
	BehindCount   int             `json:"behindCount"`
	IsBaseVersion bool            `json:"isBaseVersion"`
}

type azureDefaultBranchAPI struct {
	Count       int           `json:"count"`
	BranchValue []BranchValue `json:"value"`
}

type Headers struct {
	Token string
}

type CommitsValue struct {
	CommitID string `json:"commitId"`
	Author   struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	} `json:"author"`
	Committer struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	} `json:"committer"`
	Comment      string `json:"comment"`
	ChangeCounts struct {
		Add    int `json:"add"`
		Edit   int `json:"edit"`
		Delete int `json:"delete"`
	} `json:"changeCounts"`
	Changes []struct {
		SourceServerItem string `json:"sourceServerItem"`
		ChangeType       string `json:"changeType"`
	} `json:"changes"`
	URL       string `json:"url"`
	RemoteURL string `json:"remoteUrl"`
}

// LatestCommit returned structure
type Commit struct {
	Count        int            `json:"count"`
	CommitsValue []CommitsValue `json:"value"`
}
