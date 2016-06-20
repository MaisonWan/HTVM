package main

// 命令行解析内容
type Command struct {
	// 是否是帮助的标识
	helpFlag    bool
	// 版本标识
	versionFlag bool
	// 类路径
	cpOption    string
	// jre路径
	XjreOption  string
	// 运行的主类class
	class       string
	// Java运行参数
	args        []string
}