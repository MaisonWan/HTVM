package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"errors"
)

type ZipClass struct {
	absolutePath string
}

func newZipClass (path string) *ZipClass {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipClass{absPath}
}

// 从Zip里加载class
func (self *ZipClass) readClass(className string) ([]byte, Class, error) {
	reader, err := zip.OpenReader(self.absolutePath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.Name == className {
			rc, err := file.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipClass) String() string {
	return self.absolutePath
}