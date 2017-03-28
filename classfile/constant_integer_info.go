package classfile

type ConstantIntegerInfo struct {
	value int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	value := reader.ReadUint32()
	self.value = int32(value)
}
