package classfile

// 常量类信息
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.ReadUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.GetClassName(self.nameIndex)
}
