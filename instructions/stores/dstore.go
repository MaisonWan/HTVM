package stores

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 栈顶弹出double类型的值，然后写入临时变量区
type DSTORE struct {
	base.Index8Instruction
}

func (self *DSTORE) Execute(frame *runtime.Frame) {
	_executeDouble(frame, self.Index)
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_0) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 0)
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_1) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 1)
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_2) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 2)
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_3) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 3)
}

type DSTORE_4 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_4) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 4)
}

func _executeDouble(frame *runtime.Frame, index uint) {
	val := frame.OperateStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
