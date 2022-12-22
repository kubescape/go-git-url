package gitlabapi

import "fmt"

// ListAll list all file/dir in repo tree
func (t *Tree) ListAll() []string {
	l := []string{}

	for i := range *t {
		l = append(l, (*t)[i].Path)
	}
	return l
}

// ListAllFiles list all files in repo tree
func (t *Tree) ListAllFiles() []string {
	l := []string{}
	for i := range *t {
		if (*t)[i].Type == ObjectTypeFile {
			l = append(l, (*t)[i].Path)
		}
	}
	return l
}

// ListAllDirs list all directories in repo tree
func (t *Tree) ListAllDirs() []string {
	l := []string{}
	for i := range *t {
		if (*t)[i].Type == ObjectTypeDir {
			l = append(l, (*t)[i].Path)
		}
	}
	return l
}

// ToMap convert headers to map[string]string
func (h *Headers) ToMap() map[string]string {
	m := make(map[string]string)
	if h.Token != "" {
		m["Authorization"] = fmt.Sprintf("Bearer %s", h.Token)
	}
	return m
}
