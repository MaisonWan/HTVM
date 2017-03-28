package classpath

import (
	"HTVM/util"
	"os"
	"path/filepath"
)

// 类的路径，jvm加载启动类库，扩展类库，用户程序类库
type Classpath struct {
	bootClasspath Path
	extClasspath  Path
	userClasspath Path
}

// 解析启动时候java类的路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// 解析启动类和扩展类路径
func (cp *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClasspath = NewWailcardPath(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	cp.extClasspath = NewWailcardPath(jreExtPath)
}

// 解析用户classpath
func (cp *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClasspath = NewPath(cpOption)
}

// 依次从启动类路径，扩展类路径和用户类路径中搜索class文件
func (cp *Classpath) ReadClass(className string) ([]byte, Path, error) {
	className = className + ".class"
	if data, class, err := cp.bootClasspath.readClass(className); err == nil {
		return data, class, err
	}
	if data, class, err := cp.extClasspath.readClass(className); err == nil {
		return data, class, err
	}
	return cp.userClasspath.readClass(className)
}

func (cp *Classpath) String() string {
	return cp.userClasspath.String()
}

// 得到参数中的Xjre参数
func getJreDir(jreOption string) string {
	// 判断参数中的路径，优先级较高
	if jreOption != "" && util.Exists(jreOption) {
		return jreOption
	}
	if util.Exists("./jre") {
		return "./jre"
	}
	// 从系统中环境变量读取
	if javahome := os.Getenv("JAVA_HOME"); javahome != "" {
		return filepath.Join(javahome, "jre")
	}
	panic("can not found jre folder")
}