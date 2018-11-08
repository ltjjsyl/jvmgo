package classfile

type UnprasedAttribute struct{
	name string
	length uint32
	info [] byte
}

func(self *UnprasedAttribute) readInfo(reader *ClassReader){
	self.info = reader.readBytes(self.length)
}