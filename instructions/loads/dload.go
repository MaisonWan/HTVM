package load

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 从局部变量区加载一个int类型
type DLOAD struct {
	base.Index8Instruction
}

func (self *DLOAD) Execute(frame *runtime.Frame) {
	_executeDouble(frame, self.Index)
}

// 载入第0个位置的变量
type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_0) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 0)
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_1) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 1)
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_2) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 2)
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_3) Execute(frame *runtime.Frame) {
	_executeDouble(frame, 3)
}

func _executeDouble(frame *runtime.Frame, index uint) {
	i := frame.LocalVars().GetDouble(index)
	frame.OperateStack().PushDouble(i)
}

