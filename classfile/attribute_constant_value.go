package classfile

// 静态变量属性
type AttributeConstantValue struct {
	valueIndex uint16
}

func (self *AttributeConstantValue) readInfo(reader *ClassReader) {
	self.valueIndex = reader.readUint16()
}

// 返回常量索引的下标
func (self *AttributeConstantValue) ConstantValueIndex() uint16 {
	return self.valueIndex
}

