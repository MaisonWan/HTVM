package classfile

// 源代码的属性
type AttributeSourceFile struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *AttributeSourceFile) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.ReadUint16()
}

func (self *AttributeSourceFile) FileName() string {
	return self.cp.GetUtf8(self.sourceFileIndex)
}

