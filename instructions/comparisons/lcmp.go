package comparisons

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// compare long
type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}