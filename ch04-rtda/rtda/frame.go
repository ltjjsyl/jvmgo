package rtda

type Frame struct {
	lower        *Frame
	localVals    LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVals:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) LocalVars() LocalVars {
	return self.localVals
}
