package classfile


type ConstantMemberInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.ReadUint16()
	self.nameAndTypeIndex = reader.ReadUint16()
}

func (self *ConstantMemberInfo) ClassName() string {
	return self.cp.GetClassName(self.classIndex)
}

func (self *ConstantMemberInfo) NameAndDescriptor() (string, string) {
	return self.cp.GetNameAndType(self.nameAndTypeIndex)
}

type ConstantFieldInfo struct {
	ConstantMemberInfo
}

type ConstantMethodInfo struct {
	ConstantMemberInfo
}

type ConstantInterfaceMethodInfo struct {
	ConstantMemberInfo
}

