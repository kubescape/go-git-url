package gitlabapi

import (
	"time"
)

type ObjectType string

const (
	ObjectTypeDir  ObjectType = "tree"
	ObjectTypeFile ObjectType = "blob"
)

type InnerTree struct {
	ID   string     `json:"id"`
	Name string 	`json:"name"`
	Type ObjectType `json:"type"`
	Path string     `json:"path"`
	Mode string     `json:"mode"`
}

type Tree []InnerTree

type gitlabDefaultBranchAPI []struct {
	Name 				string 		`json:"name"`
	Commit 				Commit 		`json:"commit"`
	Merged 				bool 		`json:"merged"`
	Protected 			bool 		`json:"protected"`
	DevelopersCanPush	bool		`json:"developers_can_push"`
	DevelopersCanMerge	bool		`json:"developers_can_merge"`
	CanPush 			bool 		`json:"can_push"`
	DefaultBranch 		bool 		`json:"default"`
	WebURL 				string		`json:"web_url"`
}

type Headers struct {
	Token string
}

// LatestCommit returned structure
type Commit struct {
	ID          	string          `json:"id"`
	ShortID			string			`json:"short_id"`
	CreatedAt		time.Time		`json:"created_at"`
	Title 			string 			`json:"title"`
	Message 		string 			`json:"message"`
	AuthorName		string			`json:"author_name"`
	AuthorEmail		string			`json:"author_email"`
	AuthoredDate 	time.Time		`json:"authored_date"`
	CommitterName	string			`json:"committer_name"`
	CommitterEmail	string			`json:"committer_email"`
	CommitterDate	time.Time		`json:"committed_date"`
	WebURL 			string			`json:"web_url"`
	ParentIDS       []string		`json:"parent_ids"`
}