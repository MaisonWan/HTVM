package classfile

type ConstantUtf8Info struct {
	value string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.value = decodeMUtf8(bytes)
}

// 解析mutf8的形式
func decodeMUtf8(bytes []byte) string {
	return string(bytes)
}