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

// 读取信息
func (self *AttributeCode) readInfo(reader *ClassReader) {
	self.maxStack = reader.ReadUint16()
	self.maxLocals = reader.ReadUint16()
	codeLength := reader.ReadUint32()
	self.code = reader.ReadBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func (self *AttributeCode) MaxStack() uint16 {
	return self.maxStack
}

func (self *AttributeCode) MaxLocals() uint16 {
	return self.maxLocals
}

func (self *AttributeCode) Code() []byte {
	return self.code
}

// 异常表
type ExceptionTableInfo struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

// 读取异常表
func readExceptionTable(reader *ClassReader) []*ExceptionTableInfo {
	tableLength := reader.ReadUint16()
	exceptionTable := make([]*ExceptionTableInfo, tableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableInfo{
			startPc:  reader.ReadUint16(),
			endPc:    reader.ReadUint16(),
			handlerPc:reader.ReadUint16(),
			catchType:reader.ReadUint16(),
		}
	}
	return exceptionTable
}

func (self *ExceptionTableInfo) StartPC() uint16 {
	return self.startPc
}

func (self *ExceptionTableInfo) EndPC() uint16 {
	return self.endPc
}

func (self *ExceptionTableInfo) HandlerPC() uint16 {
	return self.handlerPc
}

func (self *ExceptionTableInfo) CatchType() uint16 {
	return self.catchType
}