package stores

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 栈顶弹出int类型的值，然后写入临时变量区
type ISOTRE struct {
	base.Index8Instruction
}

func (self *ISOTRE) Execute(frame *runtime.Frame) {
	_executeInt(frame, self.Index)
}

type ISOTRE_0 struct {
	base.NoOperandsInstruction
}

func (self *ISOTRE_0) Execute(frame *runtime.Frame) {
	_executeInt(frame, 0)
}

type ISOTRE_1 struct {
	base.NoOperandsInstruction
}

func (self *ISOTRE_1) Execute(frame *runtime.Frame) {
	_executeInt(frame, 1)
}

type ISOTRE_2 struct {
	base.NoOperandsInstruction
}

func (self *ISOTRE_2) Execute(frame *runtime.Frame) {
	_executeInt(frame, 2)
}

type ISOTRE_3 struct {
	base.NoOperandsInstruction
}

func (self *ISOTRE_3) Execute(frame *runtime.Frame) {
	_executeInt(frame, 3)
}

type ISOTRE_4 struct {
	base.NoOperandsInstruction
}

func (self *ISOTRE_4) Execute(frame *runtime.Frame) {
	_executeInt(frame, 4)
}

func _executeInt(frame *runtime.Frame, index uint) {
	val := frame.OperateStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}