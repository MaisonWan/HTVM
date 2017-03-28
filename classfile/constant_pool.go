package classfile


// 常量池
type ConstantPool []ConstantInfo

func ReadConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.ReadUint16())
	cp := make([]ConstantInfo, cpCount)

	// 索引从1开始
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 这两个类型8个字节，i是4个字节，需要多走一步
		}
	}
	return cp
}

func (self ConstantPool) GetConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invaild constant pool index!")
}

func (self ConstantPool) GetNameAndType(index uint16) (string, string) {
	ntInfo := self.GetConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.GetUtf8(ntInfo.nameIndex)
	descriptor := self.GetUtf8(ntInfo.descriptorIndex)
	return name, descriptor
}

// 得到类名
func (self ConstantPool) GetClassName(index uint16) string {
	classInfo := self.GetConstantInfo(index).(*ConstantClassInfo)
	return self.GetUtf8(classInfo.nameIndex)
}

// 得到utf8字符集
func (self ConstantPool) GetUtf8(index uint16) string {
	utf8Info := self.GetConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.value
}

