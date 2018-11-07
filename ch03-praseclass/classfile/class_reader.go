package classfile

import "encoding/binary"
import ""

type ClassReader struct{
	data []byte
}

func (self *ClassReader) readUint8() uint8{  //u1
	val := self.data[0]
	self.data = self.date[1:]
	return val
}

func (self *ClassReader) readUint16() uint16{ //u2
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.date[2:]
	return val
}

func (self *ClassReader) readUint32() uint32{
	val := binary.BigEndian,Uint32(self.data)
	self.data = self.date[4:]
	return val
}

func (self *ClassReader) readUint64() uint64{ //u4
	val := binary.BigEndian,Uint64(self.data)
	self.data = self.date[8:]
	return val
}

func (self *ClassReader) readUint16s() uint16{ //读取uint16表
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s{
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(n uint32) []byte{
	bytes := self.data[:n]
	self.data = self.date[n:]
	return bytes
}

f