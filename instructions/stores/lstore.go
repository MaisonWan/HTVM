package stores

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 栈顶弹出long类型的值，然后写入临时变量区
type LSTORE struct {
	base.Index8Instruction
}

func (self *LSTORE) Execute(frame *runtime.Frame) {
	_executeLong(frame, self.Index)
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_0) Execute(frame *runtime.Frame) {
	_executeLong(frame, 0)
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_1) Execute(frame *runtime.Frame) {
	_executeLong(frame, 1)
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_2) Execute(frame *runtime.Frame) {
	_executeLong(frame, 2)
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_3) Execute(frame *runtime.Frame) {
	_executeLong(frame, 3)
}

type LSTORE_4 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_4) Execute(frame *runtime.Frame) {
	_executeLong(frame, 4)
}

func _executeLong(frame *runtime.Frame, index uint) {
	val := frame.OperateStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}