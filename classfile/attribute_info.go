package classfile

// 属性
type AttributeInfo interface {
	// 读取信息
	readInfo(reader *ClassReader)
}

// 从数据流中读取一个属性值
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	nameIndex := reader.ReadUint16()
	attrName := cp.GetUtf8(nameIndex)
	length := reader.ReadUint32()
	attrInfo := newAttributeInfo(attrName, length, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

// 读取多个属性
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	count := reader.ReadUint16()
	attributes := make([]AttributeInfo, count)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// 创建一个新的属性
func newAttributeInfo(attrName string, length uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &AttributeCode{cp:cp}
	case "ConstantValue":
		return &AttributeConstantValue{}
	case "Deprecated":
		return &AttributeDeprecated{}
	case "Exceptions":
		return &AttributeExceptions{}
	case "LineNumberTable":
		return &AttributeLineNumberTable{}
	case "LocalVariableTable":
		return &AttributeLocalVariable{}
	case "SourceFile":
		return &AttributeSourceFile{cp:cp}
	case "Synthetic":
		return &AttributeSynthetic{}
	default:
		return &UnparsedAttribute{attrName, length, nil}
	}
}