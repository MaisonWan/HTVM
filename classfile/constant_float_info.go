package classfile

import "math"

type ConstantFloatInfo struct {
	value float32
}

// 读取一个float
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	value := reader.readUint32()
	self.value = math.Float32frombits(value)
}
