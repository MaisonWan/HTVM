package math

import (
"HTVM/instructions/base"
"HTVM/runtime"
)

type IAND struct {
	base.NoOperandsInstruction
}

func (self *IAND) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	stack.PushInt(v1 & v2)
}

type LAND struct {
	base.NoOperandsInstruction
}

func (self *LAND) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	stack.PushLong(v1 & v2)
}