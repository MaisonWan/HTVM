package math

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// double加法
type DADD struct {
	base.NoOperandsInstruction
}

func (self *DADD) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	stack.PushDouble(v1 + v2)
}

// float加法
type FADD struct {
	base.NoOperandsInstruction
}

func (self *FADD) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	stack.PushFloat(v1 + v2)
}

// int加法
type IADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	stack.PushInt(v1 + v2)
}

// long加法
type LADD struct {
	base.NoOperandsInstruction
}

func (self *LADD) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	stack.PushLong(v1 + v2)
}
