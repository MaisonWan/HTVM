package conversions

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 强制类型转换，double转化为float
type D2F struct {
	base.NoOperandsInstruction
}

func (self *D2F) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopDouble()
	stack.PushFloat(float32(v))
}

type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2I) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopDouble()
	stack.PushInt(int32(v))
}

type D2L struct {
	base.NoOperandsInstruction
}

func (self *D2L) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopDouble()
	stack.PushLong(int64(v))
}