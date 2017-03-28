package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"errors"
)

type ZipPath struct {
	absolutePath string
}

// jar包方式的路径
func newZipPath(path string) *ZipPath {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipPath{absPath}
}

// 从Zip里加载class
func (zp *ZipPath) readClass(className string) ([]byte, Path, error) {
	reader, err := zip.OpenReader(zp.absolutePath)
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
			return data, zp, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (zp *ZipPath) String() string {
	return zp.absolutePath
}