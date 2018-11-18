package control

import "jvmgo/ch05-instructions/instructions/base"
import "jvmgo/ch05-instructions/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
