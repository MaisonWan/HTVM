package classfile

type AttributeLocalVariable struct {
	localVariableTable []*LocalVariableInfo
}

type LocalVariableInfo struct {
	startPc    uint16
	lineNumber uint16
}

// 读取信息
func (self *AttributeLocalVariable) readInfo(reader *ClassReader) {
	length := reader.ReadUint16()
	self.localVariableTable = make([]*LocalVariableInfo, length)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableInfo{
			startPc:reader.ReadUint16(),
			lineNumber:reader.ReadUint16(),
		}
	}
}
