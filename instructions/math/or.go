package math

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

type IOR struct {
	base.NoOperandsInstruction
}

func (self *IOR) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	stack.PushInt(v1 | v2)
}

type LOR struct {
	base.NoOperandsInstruction
}

func (self *LOR) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	stack.PushLong(v1 | v2)
}