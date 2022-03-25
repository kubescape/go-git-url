package urlgit

type IGitURL interface {

	// parse url
	Parse(fullURL string) error

	SetBranch(string)
	SetOwner(string)
	SetPath(string)
	SetToken(string)

	GetBranch() string
	GetOwner() string
	GetPath() string
	GetToken() string

	// ListFiles list all files in path with the desired extension. if empty will list all (including directories)
	ListFiles(exctensions []string) ([]string, error)
}
