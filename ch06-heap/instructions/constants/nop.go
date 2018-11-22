package constants

import (
	"jvmgo/ch06-heap/instructions/base"
	"jvmgo/ch06-heap/rtda"
)

//TODO
type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
}
