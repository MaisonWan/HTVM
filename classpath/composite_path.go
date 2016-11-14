package classpath

import (
	"errors"
	"strings"
)

type CompositePath []Path

// 根据路径分隔符(;)，分割开来，返回数组
func NewCompositePath(pathList string) CompositePath {
	compositeClass := []Path{}
	var paths = strings.Split(pathList, pathListSeparator)
	for _, path := range paths {
		class := NewPath(path)
		compositeClass = append(compositeClass, class)
	}
	return compositeClass
}

// 读取每个子路径然后返回
func (self CompositePath) readClass(className string) ([]byte, Path, error) {
	for _, class := range self {
		data, from, err := class.readClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

// 将每个子路径组合在一起
func (self CompositePath) String() string {
	paths := make([]string, len(self))
	for i, class := range self {
		paths[i] = class.String()
	}
	return strings.Join(paths, pathListSeparator)
}