package extend

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 与goto不同的是，跳转的是4字节
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader)  {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame runtime.Frame) {
	base.Jump(frame, self.offset)
}
