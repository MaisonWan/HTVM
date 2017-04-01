package control

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

type LOOKUP_SWITCH struct {
	defaultOffset int
	caseCount     int
	mapOffset     []int
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.caseCount = reader.ReadInt32()
	// map的结构是key后面跟着value，成对出现
	self.mapOffset = reader.ReadInt32s(self.caseCount * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame runtime.Frame) {
	key := frame.OperateStack().PopInt()
	for i := 0; i < self.caseCount * 2; i += 2 {
		if self.mapOffset[i] == key {
			offset := self.mapOffset[i + 1]
			base.Jump(frame, offset)
			return
		}
	}
	base.Jump(frame, self.defaultOffset)
}
