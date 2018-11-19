package constants

import (
	"jvmgo/ch05-instructions/instructions/base"
	"jvmgo/ch05-instructions/rtda"
)

//TODO
type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
}
