package stores

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 栈顶弹出float类型的值，然后写入临时变量区
type ASOTRE struct {
	base.Index8Instruction
}

func (self *ASOTRE) Execute(frame *runtime.Frame) {
	_executeRef(frame, self.Index)
}

type ASOTRE_0 struct {
	base.NoOperandsInstruction
}

func (self *ASOTRE_0) Execute(frame *runtime.Frame) {
	_executeRef(frame, 0)
}

type ASOTRE_1 struct {
	base.NoOperandsInstruction
}

func (self *ASOTRE_1) Execute(frame *runtime.Frame) {
	_executeRef(frame, 1)
}

type ASOTRE_2 struct {
	base.NoOperandsInstruction
}

func (self *ASOTRE_2) Execute(frame *runtime.Frame) {
	_executeRef(frame, 2)
}

type ASOTRE_3 struct {
	base.NoOperandsInstruction
}

func (self *ASOTRE_3) Execute(frame *runtime.Frame) {
	_executeRef(frame, 3)
}

type ASOTRE_4 struct {
	base.NoOperandsInstruction
}

func (self *ASOTRE_4) Execute(frame *runtime.Frame) {
	_executeRef(frame, 4)
}

func _executeRef(frame *runtime.Frame, index uint) {
	val := frame.OperateStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}


