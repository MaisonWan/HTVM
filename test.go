package main

import (
	"fmt"
	"HTVM/classpath"
	"strings"
	"HTVM/classfile"
	"strconv"
)

var startJVM = func(cmd *Command) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath: %s class: %s args: %v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	classFile, _ := classfile.Parser(classData)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	//fmt.Printf("class data: %v\n", classData)
	fmt.Println("MajorVersion: ", classFile.MajorVersion())
	fmt.Println("ClassName: ", classFile.ClassName())
	fmt.Println("AccessFlags: 0x", strconv.FormatInt(int64(classFile.AccessFlags()), 16))
	fmt.Println("InterfaceNames: ", classFile.InterfaceNames())
	fmt.Println("Fields: ", classFile.InterfaceNames())
	//var info *classfile.MemberInfo
	for i, info := range classFile.Fields() {
		fmt.Println(i, info.Name(), info.Descriptor())
	}
}

func main() {
	cmd := ParseCommand()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		PrintUsage()
	} else {
		startJVM(cmd)
	}
}
