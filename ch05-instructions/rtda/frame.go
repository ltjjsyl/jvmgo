package rtda

type Frame struct {
	lower        *Frame
	localVals    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
}

func NewFrame(maxLocals, maxStack uint16) *Frame {
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
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}
func (self *Frame) Thread() *Thread {
	return self.thread
}

func newFrame(thread *Thread, maxLocals, maxStack uint16) *Frame {
	return &Frame{
		thread:       thread,
		localVals:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
