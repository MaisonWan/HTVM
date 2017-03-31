package stack

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 交换栈顶两个元素
type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
