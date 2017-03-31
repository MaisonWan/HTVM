package math

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

type DDIV struct {
	base.NoOperandsInstruction
}

func (self *DDIV) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushDouble(v1 / v2)
}

type FDIV struct {
	base.NoOperandsInstruction
}

func (self *FDIV) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushFloat(v1 / v2)
}

type IDIV struct {
	base.NoOperandsInstruction
}

func (self *IDIV) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushInt(v1 / v2)
}

type LDIV struct {
	base.NoOperandsInstruction
}

func (self *LDIV) Execute(frame *runtime.Frame) {
	stack := frame.OperateStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushLong(v1 / v2)
}
