package heap

import "jvmgo/ch06-heap/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConatantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
