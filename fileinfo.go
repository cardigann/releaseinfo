package releaseinfo

import "errors"

type FileInfo struct {
	Path      string
	Name      string
	Extension string
	Directory struct {
		Name string
	}
}

func ParseFileInfo(f string) (FileInfo, error) {
	return FileInfo{}, errors.New("Not implemented")
}
