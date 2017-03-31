package comparisons

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 对比flaot
type FCMPG struct {
	base.NoOperandsInstruction
}

func (self *FCMPG) Execute(frame *runtime.Frame) {
	_fcmp(frame, true)
}

// 对比结果有Not a number
type FCMPL struct {
	base.NoOperandsInstruction
}

func (self *FCMPL) Execute(frame *runtime.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *runtime.Frame, flag bool) {
	stack := frame.OperateStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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
