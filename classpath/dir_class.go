package classpath

import (
	"path/filepath"
	"io/ioutil"
)

// 目录形式的类路径
type DirClass struct {
	adsoluteDir string
}

func newDirClass(path string) *DirClass {
	absoluteDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirClass{absoluteDir}
}

// 加载类数据
func (self *DirClass) readClass(className string) ([]byte, Class, error) {
	fileName := filepath.Join(self.adsoluteDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

// 返回路径
func (self *DirClass) String() string {
	return self.adsoluteDir
}