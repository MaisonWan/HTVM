package comparisons

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

//对比条件if equal
type IF_ICMP_EQ struct {
	base.BranchInstruction
}

func (self *IF_ICMP_EQ) Execute(frame *runtime.Frame) {
	v2 := frame.OperateStack().PopInt()
	v1 := frame.OperateStack().PopInt()
	if v1 == v2 {
		// 如果相等
		base.Jump(frame, self.Offset)
	}
}

// if not equal
type IF_ICMP_NE struct {
	base.BranchInstruction
}

func (self *IF_ICMP_NE) Execute(frame *runtime.Frame) {
	v2 := frame.OperateStack().PopInt()
	v1 := frame.OperateStack().PopInt()
	if v1 != v2 {
		// 如果不相等
		base.Jump(frame, self.Offset)
	}
}

// if less than
type IF_ICMP_LT struct {
	base.BranchInstruction
}

func (self *IF_ICMP_LT) Execute(frame *runtime.Frame) {
	v2 := frame.OperateStack().PopInt()
	v1 := frame.OperateStack().PopInt()
	if v1 < v2 {
		base.Jump(frame, self.Offset)
	}
}

// if less equal
type IF_ICMP_LE struct {
	base.BranchInstruction
}

func (self *IF_ICMP_LE) Execute(frame *runtime.Frame) {
	v2 := frame.OperateStack().PopInt()
	v1 := frame.OperateStack().PopInt()
	if v1 <= v2 {
		base.Jump(frame, self.Offset)
	}
}

// if great
type IF_ICMP_GT struct {
	base.BranchInstruction
}

func (self *IF_ICMP_GT) Execute(frame *runtime.Frame) {
	v2 := frame.OperateStack().PopInt()
	v1 := frame.OperateStack().PopInt()
	if v1 > v2 {
		base.Jump(frame, self.Offset)
	}
}

// if great equal
type IF_ICMP_GE struct {
	base.BranchInstruction
}

func (self *IF_ICMP_GE) Execute(frame *runtime.Frame) {
	v2 := frame.OperateStack().PopInt()
	v1 := frame.OperateStack().PopInt()
	if v1 >= v2 {
		base.Jump(frame, self.Offset)
	}
}
