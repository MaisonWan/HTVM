package conversions

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

type I2F struct {
	base.NoOperandsInstruction
}

func (self *I2F) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopInt()
	stack.PushFloat(float32(v))
}

type I2D struct {
	base.NoOperandsInstruction
}

func (self *I2D) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopInt()
	stack.PushDouble(float64(v))
}

type I2L struct {
	base.NoOperandsInstruction
}

func (self *I2L) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopInt()
	stack.PushLong(int64(v))
}
