package classfile

// 异常的属性
type AttributeExceptions struct {
	exceptionIndexTable []uint16
}

func (self *AttributeExceptions) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *AttributeExceptions) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}

