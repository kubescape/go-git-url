package githubapi

import "fmt"

// ListAll list all file/dir in repo tree
func (t *Tree) ListAll() []string {
	l := []string{}
	for i := range t.InnerTrees {
		l = append(l, t.InnerTrees[i].Path)
	}
	return l
}

// ListAllFiles list all files in repo tree
func (t *Tree) ListAllFiles() []string {
	l := []string{}
	for i := range t.InnerTrees {
		if t.InnerTrees[i].Type == ObjectTypeFile {
			l = append(l, t.InnerTrees[i].Path)
		}
	}
	return l
}

// ListAllDirs list all directories in repo tree
func (t *Tree) ListAllDirs() []string {
	l := []string{}
	for i := range t.InnerTrees {
		if t.InnerTrees[i].Type == ObjectTypeDir {
			l = append(l, t.InnerTrees[i].Path)
		}
	}
	return l
}

// ToMap convert headser to map[string]string
func (h *Headres) ToMap() map[string]string {
	m := make(map[string]string)
	if h.Token != "" {
		m["Authorization"] = fmt.Sprintf("token %s", h.Token)
	}
	return m
}
