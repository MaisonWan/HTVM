package classfile


type ConstantLongInfo struct {
	value int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	value := reader.ReadUint64()
	self.value = int64(value)
}
