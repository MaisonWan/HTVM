package classfile

import "math"

type ConstantDoubleInfo struct {
	value float64
}

// 读取8字节的double
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	value := reader.readUint64()
	self.value = math.Float64frombits(value)
}
