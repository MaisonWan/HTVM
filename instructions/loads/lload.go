package load

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 从局部变量区加载一个int类型
type LLOAD struct {
	base.Index8Instruction
}

func (self *LLOAD) Execute(frame *runtime.Frame) {
	_executeLong(frame, self.Index)
}

// 载入第0个位置的变量
type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_0) Execute(frame *runtime.Frame) {
	_executeLong(frame, 0)
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_1) Execute(frame *runtime.Frame) {
	_executeLong(frame, 1)
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_2) Execute(frame *runtime.Frame) {
	_executeLong(frame, 2)
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_3) Execute(frame *runtime.Frame) {
	_executeLong(frame, 3)
}

func _executeLong(frame *runtime.Frame, index uint) {
	i := frame.LocalVars().GetLong(index)
	frame.OperateStack().PushLong(i)
}
