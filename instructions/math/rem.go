package math

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
	"math"
)

// double求余运算
type DREM struct {
	base.NoOperandsInstruction
}

func (self *DREM) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	r := math.Mod(v2, v1)
	stack.PushDouble(r)
}

// float求余运算
type FREM struct {
	base.NoOperandsInstruction
}

func (self *FREM) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	r := math.Mod(float64(v1), float64(v2))
	stack.PushFloat(float32(r))
}

// int求余运算
type IREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushInt(v1 % v2)
}

// long求余运算
type LREM struct {
	base.NoOperandsInstruction
}

func (self *LREM) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushLong(v1 % v2)
}
