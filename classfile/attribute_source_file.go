package classfile

// 源代码的属性
type AttributeSourceFile struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *AttributeSourceFile) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *AttributeSourceFile) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}

