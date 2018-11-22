package main

import (
	"fmt"
	"jvmgo/ch06-heap/classfile"
	"jvmgo/ch06-heap/classpath"
	"jvmgo/ch06-heap/rtda/heap"
	"strings"
)

// ./ch04-rtda ddd
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	classLoader := heap.NewClassLoader(cp)
	className := strings.Replace(cmd.class, ".", "/", -1)
	fmt.Printf("className:%v\n", className)
	mainClass := classLoader.LoadClass(className)
	fmt.Printf(cmd.class)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("classData:%v\n", classData)
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}
