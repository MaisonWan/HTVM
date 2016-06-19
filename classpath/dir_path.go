package classpath

import (
	"path/filepath"
	"io/ioutil"
)

// 目录形式的类路径
type DirPath struct {
	adsoluteDir string
}

func newDirPath(path string) *DirPath {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirPath{absDir}
}

// 加载类数据
func (self *DirPath) readClass(className string) ([]byte, Path, error) {
	fileName := filepath.Join(self.adsoluteDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

// 返回路径
func (self *DirPath) String() string {
	return self.adsoluteDir
}