package heap

import (
	"fmt"
	"jvmgo/ch06-heap/classfile"
)

type Constant interface{}
type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	counts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, counts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			counts[i] = intInfo.Value() // int32

		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			counts[i] = floatInfo.Value() // float32

		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			counts[i] = longInfo.Value() // int64
			i++

		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			counts[i] = doubleInfo.Value() // float64
			i++

		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			counts[i] = stringInfo.String() // string

		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			counts[i] = newClassRef(rtCp, classInfo) // class

		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			counts[i] = newFieldRef(rtCp, fieldrefInfo) // fieldref

		case *classfile.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			counts[i] = newMethodRef(rtCp, methodrefInfo) // methodref

		case *classfile.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			counts[i] = newInterfaceMethodRef(rtCp, methodrefInfo) // methodref
		}
	}
	return rtCp
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
