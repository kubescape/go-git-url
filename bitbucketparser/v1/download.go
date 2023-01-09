package v1

import "fmt"

func (gl *BitBucketURL) DownloadAllFiles() (map[string][]byte, map[string]error) {
	//TODO implement me
	return nil, map[string]error{"": fmt.Errorf("DownloadAllFiles is not supported")}
}

func (gl *BitBucketURL) DownloadFilesWithExtension(extensions []string) (map[string][]byte, map[string]error) {
	//TODO implement me
	return nil, map[string]error{"": fmt.Errorf("DownloadFilesWithExtension is not supported")}
}
