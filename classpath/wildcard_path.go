package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// 通配符和多路径类型一样，共用
func NewWailcardPath(path string) CompositePath {
	var length = len(path)
	baseDir := path[:length - 1]
	compositeClass := []Path{}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		var lowPath = strings.ToLower(path)
		// 如果是jar包
		if strings.HasSuffix(lowPath, ".jar") {
			jarClass := newZipPath(path)
			compositeClass = append(compositeClass, jarClass)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFunc)
	return compositeClass
}