package conversions

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 强制类型转换，float转化为double
type F2D struct {
	base.NoOperandsInstruction
}

func (self *F2D) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopFloat()
	stack.PushDouble(float64(v))
}

type F2I struct {
	base.NoOperandsInstruction
}

func (self *F2I) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopFloat()
	stack.PushInt(int32(v))
}

type F2L struct {
	base.NoOperandsInstruction
}

func (self *F2L) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopFloat()
	stack.PushLong(int64(v))
}