package rtda

import "jvmgo/ch06-heap/rtda/heap"

type Frame struct {
	lower        *Frame
	localVals    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int
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
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}
func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) Method() *heap.Method {
	return self.method
}
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVals:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}
