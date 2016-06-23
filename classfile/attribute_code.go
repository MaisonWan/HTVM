package classfile

// 代码属性，存储方法的内容
type AttributeCode struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableInfo
	attributes     []AttributeInfo
}

// 异常表
type ExceptionTableInfo struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

// 读取信息
func (self *AttributeCode) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

// 读取异常表
func readExceptionTable(reader *ClassReader) []*ExceptionTableInfo {
	tableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableInfo, tableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableInfo{
			startPc:reader.readUint16(),
			endPc:reader.readUint16(),
			handlerPc:reader.readUint16(),
			catchType:reader.readUint16(),
		}
	}
	return exceptionTable
}