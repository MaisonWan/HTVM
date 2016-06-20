package classfile

// 属性
type AttributeInfo interface {
	// 读取信息
	readInfo(reader *ClassReader)
}

// 从数据流中读取一个属性值
//func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
//
//}
