package stack

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 从栈顶弹出一个变量的指令
type POP struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	stack.PopSlot()
}

// 从栈顶弹出两个变量的指令
type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	stack.PopSlot()
	stack.PopSlot()
}