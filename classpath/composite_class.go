package classpath

import (
	"errors"
	"strings"
)

type CompositeClass []Class

func NewCompositeClass(pathList string) CompositeClass {
	compositeClass := []Class{}
	var paths = strings.Split(pathList, pathListSeparator)
	for _, path := range paths {
		class := NewClass(path)
		compositeClass = append(compositeClass, class)
	}
	return compositeClass
}

// 读取每个子路径然后返回
func (self CompositeClass) readClass(className string) ([]byte, Class, error) {
	for _, class := range self {
		data, from, err := class.readClass(className)
		if err == nil {
			return data, from , err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

// 将每个子路径组合在一起
func (self CompositeClass) String() string {
	paths := make([]string, len(self))
	for i, class := range self {
		paths[i] = class.String()
	}
	return strings.Join(paths, pathListSeparator)
}