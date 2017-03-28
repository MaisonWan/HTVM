package classpath

import (
	"path/filepath"
	"io/ioutil"
	"fmt"
)

// 目录形式的类路径
type DirPath struct {
	absoluteDir string
}

func newDirPath(path string) *DirPath {
	absDir, err := filepath.Abs(path)
	fmt.Println("newDirPath: " + absDir)
	if err != nil {
		panic(err)
	}
	return &DirPath{absDir}
}

// 加载类数据
func (dp *DirPath) readClass(className string) ([]byte, Path, error) {
	fileName := filepath.Join(dp.absoluteDir, className)
	fmt.Println("newDirPath readClass: " + fileName)
	data, err := ioutil.ReadFile(fileName)
	return data, dp, err
}

// 返回路径
func (dp *DirPath) String() string {
	return dp.absoluteDir
}