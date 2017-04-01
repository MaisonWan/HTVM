package conversions

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

type I2F struct {
	base.NoOperandsInstruction
}

func (self *I2F) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopInt()
	stack.PushFloat(float32(v))
}

type I2D struct {
	base.NoOperandsInstruction
}

func (self *I2D) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopInt()
	stack.PushDouble(float64(v))
}

type I2L struct {
	base.NoOperandsInstruction
}

func (self *I2L) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopInt()
	stack.PushLong(int64(v))
}

type I2B struct {
	base.NoOperandsInstruction
}

func (self *I2B) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopInt()
	stack.PushInt(int32(int8(v)))
}

// int to char
type I2C struct {
	base.NoOperandsInstruction
}

func (self *I2C) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopInt()
	stack.PushInt(int32(uint16(v)))
}

type I2S struct {
	base.NoOperandsInstruction
}

func (self *I2S) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v := stack.PopInt()
	stack.PushInt(int32(int16(v)))
}
