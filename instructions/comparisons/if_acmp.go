package comparisons

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 判断引用是否相等
type ACMPEQ struct {
	base.BranchInstruction
}

func (self *ACMPEQ) Execute(frame *runtime.Frame) {
	v2 := frame.OperateStack().PopRef()
	v1 := frame.OperateStack().PopRef()
	if v1 == v2 {
		base.Jump(frame, self.Offset)
	}
}

// 判断引用是否不相等
type ACMPNE struct {
	base.BranchInstruction
}

func (self *ACMPNE) Execute(frame *runtime.Frame) {
	v2 := frame.OperateStack().PopRef()
	v1 := frame.OperateStack().PopRef()
	if v1 != v2 {
		base.Jump(frame, self.Offset)
	}
}
