package classfile


const (
	CONSTANT_Utf8 = 1
	CONSTANT_Integer = 3
	CONSTANT_Float = 4
	CONSTANT_Long = 5
	CONSTANT_Double = 6
	CONSTANT_Class = 7
	CONSTANT_String = 8
	CONSTANT_Field = 9
	CONSTANT_Method = 10
	CONSTANT_InterfaceMethod = 11
	CONSTANT_NameAndType = 12
	//CONSTANT_MethodHandle = 15
	//CONSTANT_MethodType = 16
	//CONSTANT_InvokeDynamic = 18
)

// 常量信息类
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.ReadUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Class:
		return &ConstantClassInfo{cp:cp}
	case CONSTANT_Field:
		return &ConstantFieldInfo{ConstantMemberInfo{cp:cp}}
	case CONSTANT_Method:
		return &ConstantMethodInfo{ConstantMemberInfo{cp:cp}}
	case CONSTANT_InterfaceMethod:
		return &ConstantInterfaceMethodInfo{ConstantMemberInfo{cp:cp}}
	case CONSTANT_String:
		return &ConstantStringInfo{cp:cp}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}