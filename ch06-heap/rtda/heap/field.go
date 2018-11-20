package heap

import "jvmgo/ch06-heap/classfile"

type Field struct {
	ClassMember
	slotId          uint
	cosntValueIndex uint
}

func newFields(class *Class, cfFields []*classfile.MetmberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) copyAttributes(cfField *classfile.MetmberInfo) {
	if valAttr := cfField.ConstantVlaueAttribute(); valAttr != nil {
		self.cosntValueIndex = uint(valAttr.cosntantValueIndex)
	}
}
