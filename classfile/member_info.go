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
		accessFlags:reader.ReadUint16(),
		nameIndex:reader.ReadUint16(),
		descriptorIndex:reader.ReadUint16(),
		attributes: readAttributes(reader, cp),
	}
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.ReadUint16()
	memberInfos := make([]*MemberInfo, memberCount)
	for i := range memberInfos {
		memberInfos[i] = readMember(reader, cp)
	}
	return memberInfos
}

// 得到名字
func (self *MemberInfo) Name() string {
	return self.cp.GetUtf8(self.nameIndex)
}

// 得到描述
func (self *MemberInfo) Descriptor() string {
	return self.cp.GetUtf8(self.descriptorIndex)
}

// 属性列表
func (self *MemberInfo) AttributeInfo() []AttributeInfo {
	return self.attributes
}

func (self *MemberInfo) CodeAttibutes() *AttributeCode {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *AttributeCode:
			return attrInfo.(*AttributeCode)
		}
	}
	return nil
}