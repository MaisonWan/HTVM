package comparisons

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

//对比条件if equal
type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *runtime.Frame) {
	value := frame.OperateStack().PopInt()
	if value == 0 {
		// 如果相等
		base.Jump(frame, self.Offset)
	}
}

// if not equal
type IFNE struct {
	base.BranchInstruction
}

func (self *IFNE) Execute(frame *runtime.Frame) {
	value := frame.OperateStack().PopInt()
	if value != 0 {
		// 如果不相等
		base.Jump(frame, self.Offset)
	}
}

// if less than
type IFLT struct {
	base.BranchInstruction
}

func (self *IFLT) Execute(frame *runtime.Frame) {
	value := frame.OperateStack().PopInt()
	if value == -1 {
		base.Jump(frame, self.Offset)
	}
}

// if less equal
type IFLE struct {
	base.BranchInstruction
}

func (self *IFLE) Execute(frame *runtime.Frame) {
	value := frame.OperateStack().PopInt()
	if value <= 0 {
		base.Jump(frame, self.Offset)
	}
}

// if great
type IFGT struct {
	base.BranchInstruction
}

func (self *IFGT) Execute(frame *runtime.Frame) {
	value := frame.OperateStack().PopInt()
	if value > 0 {
		base.Jump(frame, self.Offset)
	}
}

// if great equal
type IFGE struct {
	base.BranchInstruction
}

func (self *IFGE) Execute(frame *runtime.Frame) {
	value := frame.OperateStack().PopInt()
	if value >= 0 {
		base.Jump(frame, self.Offset)
	}
}
