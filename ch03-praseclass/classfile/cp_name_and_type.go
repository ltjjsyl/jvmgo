package classfile

type ConstantClassInfo struct{
	nameIndex uint16
	descriptorIndex uint16
}

func (self * ConstantClassInfo) readInfo(reader *ClassReader){
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}


