package stack

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 复制栈顶元素，再次添加到栈顶
type DUP struct {
	base.NoOperandsInstruction
}

func (self *DUP) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// 复制栈顶元素，然后插入到两个元素的下面
type DUP_X1 struct {
	base.NoOperandsInstruction
}

func (self *DUP_X1) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// 复制栈顶元素，然后插入到三个元素的下面
type DUP_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP_X2) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// 复制栈顶两个元素
type DUP2 struct {
	base.NoOperandsInstruction
}

func (self *DUP2) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// 复制栈顶两个元素，插入到两、三个元素下面
type DUP2_X1 struct {
	base.NoOperandsInstruction
}

func (self *DUP2_X1) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// 复制栈顶两个元素，插入到两、三，四个元素下面
type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP2_X2) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}