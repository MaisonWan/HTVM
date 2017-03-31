package stores

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 栈顶弹出long类型的值，然后写入临时变量区
type LSOTRE struct {
	base.Index8Instruction
}

func (self *LSOTRE) Execute(frame *runtime.Frame) {
	_executeLong(frame, self.Index)
}

type LSOTRE_0 struct {
	base.NoOperandsInstruction
}

func (self *LSOTRE_0) Execute(frame *runtime.Frame) {
	_executeLong(frame, 0)
}

type LSOTRE_1 struct {
	base.NoOperandsInstruction
}

func (self *LSOTRE_1) Execute(frame *runtime.Frame) {
	_executeLong(frame, 1)
}

type LSOTRE_2 struct {
	base.NoOperandsInstruction
}

func (self *LSOTRE_2) Execute(frame *runtime.Frame) {
	_executeLong(frame, 2)
}

type LSOTRE_3 struct {
	base.NoOperandsInstruction
}

func (self *LSOTRE_3) Execute(frame *runtime.Frame) {
	_executeLong(frame, 3)
}

type LSOTRE_4 struct {
	base.NoOperandsInstruction
}

func (self *LSOTRE_4) Execute(frame *runtime.Frame) {
	_executeLong(frame, 4)
}

func _executeLong(frame *runtime.Frame, index uint) {
	val := frame.OperateStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}