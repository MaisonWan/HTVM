package control

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// switch
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffset    []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	count := self.high - self.low + 1
	self.jumpOffset = reader.ReadInt32s(count)
}

func (self *TABLE_SWITCH) Execute(frame *runtime.Frame) {
	value := frame.OperateStack().PopInt()
	var offset int
	if value >= self.low && value <= self.high {
		offset = int(self.jumpOffset[value - self.low])
	} else {
		// 如果不在范围之内，执行默认值
		offset = int(self.defaultOffset)
	}
	base.Jump(frame, offset)
}

