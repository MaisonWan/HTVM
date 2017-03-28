package classfile

import (
	"fmt"
)

// class文件的结构
type ClassFile struct {
	// 魔数，用于标示是否是合法class文件
	magic        uint32
	// 小版本号
	minorVersion uint16
	// 主版本号
	majorVersion uint16
	// 常量池
	constantPool ConstantPool
	// 类访问标志, 表示类或者接口，访问public(0x21)还是private
	accessFlags  uint16
	// 当前类
	thisClass    uint16
	// 父类
	superClass   uint16
	// 接口
	interfaces   []uint16
	// 字段，可能多个
	fields       []*MemberInfo
	// 方法
	methods      []*MemberInfo
	// 属性
	attributes   []AttributeInfo
}

// 解析成ClassFile
func Parser(classData []byte) (classfile *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	classfile = &ClassFile{}
	classfile.read(cr)
	return
}

// 通过使用ClassReader解析成ClassFile
func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = ReadConstantPool(reader)
	cf.accessFlags = reader.ReadUint16()
	cf.thisClass = reader.ReadUint16()
	cf.superClass = reader.ReadUint16()
	cf.interfaces = reader.ReadUint16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)
}

// 读取魔数，并验证合法性，Java class文件的魔数是0xCAFEBABE
func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.ReadUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 读取版本号，1.2之前才有小版本号，1.2大版本号是46
// Java 6: 50, Java 7: 51, Java 8: 52
func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.ReadUint16()
	cf.majorVersion = reader.ReadUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// Getter方法
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.GetClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.GetClassName(cf.superClass)
	}
	return "" // 只有java.lang.Object没有超类
}

func (cf *ClassFile) InterfaceNames() []string {
	length := len(cf.interfaces)
	interfaceNames := make([]string, length)
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.GetClassName(cpIndex)
	}
	return interfaceNames
}