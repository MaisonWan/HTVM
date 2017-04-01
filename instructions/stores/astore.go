package stores

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 栈顶弹出float类型的值，然后写入临时变量区
type ASTORE struct {
	base.Index8Instruction
}

func (self *ASTORE) Execute(frame *runtime.Frame) {
	_executeRef(frame, self.Index)
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_0) Execute(frame *runtime.Frame) {
	_executeRef(frame, 0)
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_1) Execute(frame *runtime.Frame) {
	_executeRef(frame, 1)
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_2) Execute(frame *runtime.Frame) {
	_executeRef(frame, 2)
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_3) Execute(frame *runtime.Frame) {
	_executeRef(frame, 3)
}

type ASTORE_4 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_4) Execute(frame *runtime.Frame) {
	_executeRef(frame, 4)
}

func _executeRef(frame *runtime.Frame, index uint) {
	val := frame.OperateStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}


