package conversions

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

type L2F struct {
	base.NoOperandsInstruction
}

func (self *L2F) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopLong()
	stack.PushFloat(float32(v))
}

type L2D struct {
	base.NoOperandsInstruction
}

func (self *L2D) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopLong()
	stack.PushDouble(float64(v))
}

type L2I struct {
	base.NoOperandsInstruction
}

func (self *L2I) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopLong()
	stack.PushInt(int32(v))
}
