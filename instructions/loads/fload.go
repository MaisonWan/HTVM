package load

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 从局部变量区加载一个int类型
type FLOAD struct {
	base.Index8Instruction
}

func (self *FLOAD) Execute(frame *runtime.Frame) {
	_executeFloat(frame, self.Index)
}

// 载入第0个位置的变量
type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_0) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 0)
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_1) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 1)
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_2) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 2)
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_3) Execute(frame *runtime.Frame) {
	_executeFloat(frame, 3)
}

func _executeFloat(frame *runtime.Frame, index uint) {
	i := frame.LocalVars().GetFloat(index)
	frame.OperateStack().PushFloat(i)
}
