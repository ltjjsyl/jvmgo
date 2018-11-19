package heap

import "jvmgo/ch06-heap/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConatantPool, refInfo *classfile.ConstantFiledrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
	return ref
}
