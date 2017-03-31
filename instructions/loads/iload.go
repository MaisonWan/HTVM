package load

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 从局部变量区加载一个int类型
type ILOAD struct {
	base.Index8Instruction
}

func (self *ILOAD) Execute(frame *runtime.Frame) {
	_execute(frame, self.Index)
}

// 载入第0个位置的变量
type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_0) Execute(frame *runtime.Frame) {
	_execute(frame, 0)
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_1) Execute(frame *runtime.Frame) {
	_execute(frame, 1)
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_2) Execute(frame *runtime.Frame) {
	_execute(frame, 2)
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_3) Execute(frame *runtime.Frame) {
	_execute(frame, 3)
}

func _execute(frame *runtime.Frame, index uint) {
	i := frame.LocalVars().GetInt(index)
	frame.OperateStack().PushInt(i)
}