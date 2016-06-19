package classpath

import (
	"HTVM/util"
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Class
	extClasspath Class
	userClasspath Class
}

// 解析启动时候java类的路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// 解析启动类和扩展类路径
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = NewWailcardClass(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = NewWailcardClass(jreExtPath)
}

// 解析用户classpath
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = NewClass(cpOption)
}

// 依次从启动类路径，扩展类路径和用户类路径中搜索class文件
func (self *Classpath) ReadClass(className string) ([]byte, Class, error) {
	className = className + ".class"
	if data, class, err := self.bootClasspath.readClass(className); err == nil {
		return data, class, err
	}
	if data, class, err := self.extClasspath.readClass(className); err == nil {
		return data, class, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
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