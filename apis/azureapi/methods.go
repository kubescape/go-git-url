package azureapi

import (
	b64 "encoding/base64"
	"fmt"
)

// ListAll list all file/dir in repo tree
func (t *Tree) ListAll() []string {
	l := []string{}
	for i := range t.InnerTree {
		l = append(l, t.InnerTree[i].Path)
	}
	return l
}

// ListAllFiles list all files in repo tree
func (t *Tree) ListAllFiles() []string {
	l := []string{}
	for i := range t.InnerTree {
		if t.InnerTree[i].GitObjectType == ObjectTypeFile {
			l = append(l, t.InnerTree[i].Path)
		}
	}
	return l
}

// ListAllDirs list all directories in repo tree
func (t *Tree) ListAllDirs() []string {
	l := []string{}
	for i := range t.InnerTree {
		if t.InnerTree[i].GitObjectType == ObjectTypeDir {
			l = append(l, t.InnerTree[i].Path)
		}
	}
	return l
}

// ToMap convert headers to map[string]string
func (h *Headers) ToMap() map[string]string {
	m := make(map[string]string)

	if h.Token != "" {
		fToken := ":" + h.Token
		sEnc := b64.StdEncoding.EncodeToString([]byte(fToken))
		m["Authorization"] = fmt.Sprintf("Basic %s", sEnc)
	}
	return m
}
