package load

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 从局部变量区加载一个int类型
type ALOAD struct {
	base.Index8Instruction
}

func (self *ALOAD) Execute(frame *runtime.Frame) {
	_executeRef(frame, self.Index)
}

// 载入第0个位置的变量
type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_0) Execute(frame *runtime.Frame) {
	_executeRef(frame, 0)
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_1) Execute(frame *runtime.Frame) {
	_executeRef(frame, 1)
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_2) Execute(frame *runtime.Frame) {
	_executeRef(frame, 2)
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_3) Execute(frame *runtime.Frame) {
	_executeRef(frame, 3)
}

func _executeRef(frame *runtime.Frame, index uint) {
	i := frame.LocalVars().GetRef(index)
	frame.OperateStack().PushRef(i)
}

