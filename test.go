package main

import (
	"fmt"
	"HTVM/classpath"
	"strings"
	"HTVM/classfile"
	"strconv"
	"HTVM/runtime"
)

var startJVM = func(cmd *Command) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf(" Classpath: %s\n Main Class: %s\n Args: %v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1) // 将包名转化为本地路径
	classData, _, err := cp.ReadClass(className)
	classFile, _ := classfile.Parser(classData)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	printClassFile(classFile)
}

func printClassFile(classFile *classfile.ClassFile) {
	//fmt.Printf("class data: %v\n", classData)
	fmt.Println("MajorVersion: ", classFile.MajorVersion())
	fmt.Println("ClassName: ", classFile.ClassName())
	fmt.Println("SuperClassName: " + classFile.SuperClassName())
	fmt.Println("AccessFlags: 0x" + strconv.FormatInt(int64(classFile.AccessFlags()), 16))
	fmt.Println("InterfaceNames: ", classFile.InterfaceNames())
	fmt.Println("Fields======================")
	for i, info := range classFile.Fields() {
		fmt.Println(i, info.Name(), info.Descriptor())
	}
	fmt.Println("Methods=====================")
	for i, info := range classFile.Methods() {
		fmt.Println(i, info.Name(), info.Descriptor())
	}
}

func testFrame() {
	frame := runtime.NewFrame(512, 512)
	testLocalVars(frame.LocalVars())
	testOpertateStack(frame.OperateStack())
}

func testLocalVars(v runtime.LocalVars) {
	v.SetInt(0, 123)
	v.SetInt(1, -321)
	v.SetLong(2, 2997924580)
	v.SetLong(4, -2997924580)
	v.SetFloat(6, 3.1415926)
	v.SetFloat(7, -3.1415926)
	v.SetDouble(8, 2.71828182818281828)
	v.SetDouble(10, -2.71828182818281828)
	println(v.GetInt(0))
	println(v.GetInt(1))
	println(v.GetLong(2))
	println(v.GetLong(4))
	println(v.GetFloat(6))
	println(v.GetFloat(7))
	println(v.GetDouble(8))
	println(v.GetDouble(10))
}

func testOpertateStack(s *runtime.OperateStack) {
	s.PushInt(100)
	s.PushInt(-100)
	s.PushLong(1234567890)
	s.PushLong(-1234567890)
	s.PushFloat(3.1415926)
	s.PushDouble(2.718281828)
	s.PushRef(nil)
	println(s.PopRef())
	println(s.PopDouble())
	println(s.PopFloat())
	println(s.PopLong())
	println(s.PopLong())
	println(s.PopInt())
	println(s.PopInt())
}

func main() {
	cmd := ParseCommand()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		PrintUsage()
	} else {
		startJVM(cmd)
		//testFrame()
	}
}
