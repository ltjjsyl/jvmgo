package classfile

type ConstantValueAttribute struct{
	constantValueIndex uint16
}

func (self *ConstantValueAttribute)readInfo(reader *ClassReader){
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueAttribute() uint16{
	return self.constantValueIndex 
}