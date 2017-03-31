package math

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// double乘法
type DMUL struct {
	base.NoOperandsInstruction
}

func (self *DMUL) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	stack.PushDouble(v1 * v2)
}

// float乘法
type FMUL struct {
	base.NoOperandsInstruction
}

func (self *FMUL) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	stack.PushFloat(v1 * v2)
}

// int乘法
type IMUL struct {
	base.NoOperandsInstruction
}

func (self *IMUL) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	stack.PushInt(v1 * v2)
}

// long乘法
type LMUL struct {
	base.NoOperandsInstruction
}

func (self *LMUL) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	stack.PushLong(v1 * v2)
}
