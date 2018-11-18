package main

import "fmt"
import "jvmgo/ch04-rtda/rtda"

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
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2345235235)
	vars.SetLong(4, -455345345)
	vars.SetFloat(6, 3.545100)
	vars.SetDouble(7, -3.545100333)
	vars.SetRef(9, nil)

	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))

}
func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(1003452345123456)
	ops.PushLong(-1003452345123456)
	ops.PushFloat(333.10033)
	ops.PushDouble(32534.1003412351235)
	ops.PushRef(nil)

	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}
