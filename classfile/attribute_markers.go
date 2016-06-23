package classfile

// 标记的属性
type AttributeMarker struct {

}

// 过期的标记
type AttributeDeprecated struct {
	AttributeMarker
}

type AttributeSynthetic struct {
	AttributeMarker
}

func (self *AttributeMarker) readInfo(reader *ClassReader) {
	// 只是标记，没有数据，所以不用实现具体内容
}
