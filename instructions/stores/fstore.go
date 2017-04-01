package stores

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 栈顶弹出float类型的值，然后写入临时变量区
type FSTORE struct {
	base.Index8Instruction
}

func (self *FSTORE) Execute(frame *runtime.Frame) {
	_executeFloat(frame, self.Index)
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_0) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 0)
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_1) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 1)
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_2) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 2)
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_3) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 3)
}

type FSTORE_4 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_4) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 4)
}

func _executeFloat(frame *runtime.Frame, index uint) {
	val := frame.OperateStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
