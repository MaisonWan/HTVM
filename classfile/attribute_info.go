package classfile

// 属性
type AttributeInfo interface {
	// 读取信息
	readInfo(reader *ClassReader)
}

// 从数据流中读取一个属性值
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	nameIndex := reader.readUint16()
	attrName := cp.getUtf8(nameIndex)
	length := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, length, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

// 读取多个属性
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	count := reader.readUint16()
	attributes := make([]AttributeInfo, count)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// 创建一个新的属性
func newAttributeInfo(attrName string, length int32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
	case "ConstantValue":
	case "Deprecated":
	case "Exceptions":
	case "LineNumberTable":
	case "LocalVariableTable":
	case "SourceFile":
	case "Synthetic":
	default:
		return &UnparsedAttribute{attrName, length, nil}
	}
}