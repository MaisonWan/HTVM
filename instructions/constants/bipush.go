package constants

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// byte类型转化为int，然后压栈
type BIPUSH struct {
	value int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader)  {
	self.value = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *runtime.Frame)  {
	i := int32(self.value)
	frame.OperateStack().PushInt(i)
}

// short类型转化为int，然后压栈
type SIPUSH struct {
	value int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader)  {
	self.value = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *runtime.Frame)  {
	i := int32(self.value)
	frame.OperateStack().PushInt(i)
}