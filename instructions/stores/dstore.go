package stores

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 栈顶弹出double类型的值，然后写入临时变量区
type DSOTRE struct {
	base.Index8Instruction
}

func (self *DSOTRE) Execute(frame *runtime.Frame) {
	_executeDouble(frame, self.Index)
}

type DSOTRE_0 struct {
	base.NoOperandsInstruction
}

func (self *DSOTRE_0) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 0)
}

type DSOTRE_1 struct {
	base.NoOperandsInstruction
}

func (self *DSOTRE_1) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 1)
}

type DSOTRE_2 struct {
	base.NoOperandsInstruction
}

func (self *DSOTRE_2) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 2)
}

type DSOTRE_3 struct {
	base.NoOperandsInstruction
}

func (self *DSOTRE_3) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 3)
}

type DSOTRE_4 struct {
	base.NoOperandsInstruction
}

func (self *DSOTRE_4) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 4)
}

func _executeDouble(frame *runtime.Frame, index uint) {
	val := frame.OperateStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
