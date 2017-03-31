package math

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// double求余运算
type DSUB struct {
	base.NoOperandsInstruction
}

func (self *DSUB) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	stack.PushDouble(v1 - v2)
}

// float求余运算
type FSUB struct {
	base.NoOperandsInstruction
}

func (self *FSUB) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	stack.PushFloat(v1 - v2)
}

// int求余运算
type ISUB struct {
	base.NoOperandsInstruction
}

func (self *ISUB) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	stack.PushInt(v1 - v2)
}

// long求余运算
type LSUB struct {
	base.NoOperandsInstruction
}

func (self *LSUB) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	stack.PushLong(v1 - v2)
}
