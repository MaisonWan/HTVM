package classfile

type AttributeLineNumberTable struct {
	lineNumberTable []*LineNumberTableInfo
}

type LineNumberTableInfo struct {
	startPc    uint16
	lineNumber uint16
}

// 读取信息
func (self *AttributeLineNumberTable) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.ReadUint16()
	self.lineNumberTable = make([]*LineNumberTableInfo, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableInfo{
			startPc:reader.ReadUint16(),
			lineNumber:reader.ReadUint16(),
		}
	}
}