package base

// 字节码读取
type BytecodeReader struct {
	code []byte
	pc   int
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) ReadUint8() uint8 {
	index := self.code[self.pc]
	self.pc++
	return index
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return byte1 << 8 | byte2
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// 按照指定大小读取一个int32的数组返回
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	nums := make([]int32, n)
	for i := range nums {
		nums[i] = self.ReadInt32()
	}
	return nums
}

// 跳转字节
func (self *BytecodeReader) SkipPadding() {
	if self.pc % 4 != 0 {
		self.ReadUint8()
	}
}

func (self *BytecodeReader) PC() int {
	return self.pc
}