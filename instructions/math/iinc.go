package math

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 给局部变量表中的变量增加常量值
type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *runtime.Frame) {
	vars := frame.LocalVars()
	v := vars.GetInt(self.Index)
	v += self.Const
	vars.SetInt(self.Index, v)
}

