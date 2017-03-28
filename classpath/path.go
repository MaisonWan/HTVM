package classpath

import (
	"os"
	"strings"
)

// 路径分隔符
const pathListSeparator = string(os.PathListSeparator)

type Path interface {
	// 加载class
	readClass(className string) ([]byte, Path, error)
	// 返回字符串
	String() string
}

/**
 * 根据给定的路径，创个新的加载路径
 */
func NewPath(path string) Path {
	// 如果包含分隔符，组合路径
	if strings.Contains(path, pathListSeparator) {
		return NewCompositePath(path)
	}
	// 是否包含通配符
	if strings.HasSuffix(path, "*") {
		return NewWailcardPath(path)
	}
	// 是否包含jar或者zip
	var lowPath = strings.ToLower(path)
	if strings.HasSuffix(lowPath, ".jar") || strings.HasSuffix(lowPath, ".zip") {
		return newZipPath(path)
	}

	return newDirPath(path)
}