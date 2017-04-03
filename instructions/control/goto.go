package control

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 直接跳转
type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *runtime.Frame) {
	base.Jump(frame, self.Offset)
}

