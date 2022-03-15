package source

import (
	"io/ioutil"
	"path/filepath"
)

type Source struct {
	rawCode  string
	filePath string
}

func ReadSource(filePath string) (*Source, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	return &Source{
		filePath: absPath,
		rawCode:  string(content),
	}, nil
}

func (source *Source) GetRawCode() string {
	return source.rawCode
}

func (source *Source) GetSourcePath() string {
	return source.filePath
}
