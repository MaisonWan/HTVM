package stores

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 栈顶弹出int类型的值，然后写入临时变量区
type ISTORE struct {
	base.Index8Instruction
}

func (self *ISTORE) Execute(frame *runtime.Frame) {
	_executeInt(frame, self.Index)
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_0) Execute(frame *runtime.Frame) {
	_executeInt(frame, 0)
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_1) Execute(frame *runtime.Frame) {
	_executeInt(frame, 1)
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_2) Execute(frame *runtime.Frame) {
	_executeInt(frame, 2)
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_3) Execute(frame *runtime.Frame) {
	_executeInt(frame, 3)
}

type ISTORE_4 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_4) Execute(frame *runtime.Frame) {
	_executeInt(frame, 4)
}

func _executeInt(frame *runtime.Frame, index uint) {
	val := frame.OperateStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}