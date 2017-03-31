package stores

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 栈顶弹出float类型的值，然后写入临时变量区
type FSOTRE struct {
	base.Index8Instruction
}

func (self *FSOTRE) Execute(frame *runtime.Frame) {
	_executeFloat(frame, self.Index)
}

type FSOTRE_0 struct {
	base.NoOperandsInstruction
}

func (self *FSOTRE_0) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 0)
}

type FSOTRE_1 struct {
	base.NoOperandsInstruction
}

func (self *FSOTRE_1) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 1)
}

type FSOTRE_2 struct {
	base.NoOperandsInstruction
}

func (self *FSOTRE_2) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 2)
}

type FSOTRE_3 struct {
	base.NoOperandsInstruction
}

func (self *FSOTRE_3) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 3)
}

type FSOTRE_4 struct {
	base.NoOperandsInstruction
}

func (self *FSOTRE_4) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 4)
}

func _executeFloat(frame *runtime.Frame, index uint) {
	val := frame.OperateStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
