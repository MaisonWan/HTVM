package classfile

// 字符串类型，本身就是常量，关联常量池中
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.ReadUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.GetUtf8(self.stringIndex)
}