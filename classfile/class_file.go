package classfile

import (
	"fmt"
)

type ClassFile struct {
	// 魔数，用于标示是否是合法class文件
	magic        uint32
	// 小版本号
	minorVersion uint16
	// 主版本号
	majorVersion uint16
	// 常量池
	constantPool ConstantPool
	// 类访问标志, 表示类或者接口，访问public还是private
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
func Parse(classData []byte) (cf *ClassFile, err error) {
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
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// 通过使用ClassReader解析成ClassFile
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	//self.constantPool
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	//self.fields =
	//self.methods =
	//self.attributes =
}

// 读取魔数，并验证合法性
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 读取版本号
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// Getter方法
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return "" //self.thisClass
}

func (self *ClassFile) SuperClassName() string {
	return "" // 只有java.lang.Object没有超类
}

func (self *ClassFile) InterfaceNames() []string {
	return nil
}