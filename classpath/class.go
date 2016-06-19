package classpath

import (
	"os"
	"strings"
)

// 路径分隔符
const pathListSeparator = string(os.PathListSeparator)

type Class interface {
	// 加载class
	readClass(className string) ([]byte, Class, error)
	// 返回字符串
	String() string
}

func NewClass(path string) Class {
	// 如果包含分隔符
	if strings.Contains(path, pathListSeparator) {
		return NewCompositeClass(path)
	}
	// 是否包含通配符
	if strings.HasSuffix(path, "*") {
		return NewWailcardClass(path)
	}
	// 是否包含jar或者zip
	var lowPath = strings.ToLower(path)
	if strings.HasSuffix(lowPath, ".jar") || strings.HasSuffix(lowPath, ".zip") {
		return newZipClass(path)
	}

	return newDirClass(path)
}