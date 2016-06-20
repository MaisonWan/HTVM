package classfile

//
type MemberInfo struct {
	// 常量区
	cp              ConstantPool
	// 访问权限
	accessFlags     uint16
	// 名字索引
	nameIndex       uint16
	// 描述索引
	descriptorIndex uint16
	// 属性
	attributes      []AttributeInfo
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp: cp,
		accessFlags:reader.readUint16(),
		nameIndex:reader.readUint16(),
		descriptorIndex:reader.readUint16(),
		//attributes:reader
	}
}
