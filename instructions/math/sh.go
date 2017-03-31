package math

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// int类型左移
type ISHL struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 << (uint32(v2) & 0x1F)
	stack.PushInt(result)
}

// int算数右移
type ISHR struct {
	base.NoOperandsInstruction
}

func (self *ISHR) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 >> (uint(v2) & 0x1F)
	stack.PushInt(result)
}

type IUSHR struct {
	base.NoOperandsInstruction
}

func (self *IUSHR) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := uint8(v1) >> (uint(v2) & 0x1F)
	stack.PushInt(int32(result))
}

// long类型左移
type LSHL struct {
	base.NoOperandsInstruction
}

func (self *LSHL) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	result := v1 << (uint(v2) & 0x3F)
	stack.PushInt(result)
}

// long类型算数右移
type LSHR struct {
	base.NoOperandsInstruction
}

func (self *LSHR) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	result := v1 >> (uint(v2) & 0x3F)
	stack.PushInt(result)
}

// long类型逻辑右移
type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *LUSHR) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	result := uint32(v1) >> (uint(v2) & 0x3F)
	stack.PushInt(int32(result))
}