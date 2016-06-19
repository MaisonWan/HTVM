package main

import (
	"flag"
	"fmt"
	"os"
)

// 解析命令行
func ParseCommand() *Command {
	cmd := &Command{}

	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "Print help message.")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print help meesage")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.class, "classpath", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	jargs := flag.Args()
	if len(jargs) > 0 {
		cmd.class = jargs[0]
		cmd.args = jargs[1:]
	}
	return cmd
}

// 打印命令行说明
func PrintUsage() {
	fmt.Printf("Usage: %s [-option] class [args...]\n", os.Args[0])
}