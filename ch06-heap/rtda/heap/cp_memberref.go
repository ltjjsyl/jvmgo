package heap

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMenberrefInfo) {
	self.className = refInfo.Classfile()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}
