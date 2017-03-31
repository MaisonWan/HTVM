package comparisons

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 对比double
type DCMPG struct {
	base.NoOperandsInstruction
}

func (self *DCMPG) Execute(frame *runtime.Frame) {
	_fcmpDouble(frame, true)
}

// 对比结果有Not a number
type DCMPL struct {
	base.NoOperandsInstruction
}

func (self *DCMPL) Execute(frame *runtime.Frame) {
	_fcmpDouble(frame, false)
}

func _fcmpDouble(frame *runtime.Frame, flag bool) {
	stack := frame.OperateStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if flag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
