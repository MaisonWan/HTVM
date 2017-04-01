package extend

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

type IFNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame runtime.Frame) {
	ref := frame.OperateStack().PopRef()
	if ref == nil {
		base.Jump(frame, self.Offset)
	}
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNONNULL) Execute(frame runtime.Frame) {
	ref := frame.OperateStack().PopRef()
	if ref != nil {
		base.Jump(frame, self.Offset)
	}
}