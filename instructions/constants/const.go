package constants

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

// 把null引用压栈
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (self *ACONST_NULL) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushRef(nil)
}

// 把double类型0压栈
type DCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_0) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushDouble(0.0)
}

type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_1) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushDouble(1.0)
}

type FCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_0) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushFloat(0.0)
}

type FCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_1) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushFloat(1.0)
}

type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_2) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushFloat(2.0)
}

type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_M1) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushInt(-1)
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_0) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushInt(0)
}

type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_1) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushInt(1)
}

type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_2) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushInt(2)
}

type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_3) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushInt(3)
}

type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_4) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushInt(4)
}

type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_5) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushInt(5)
}

type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_0) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushLong(0)
}

type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_1) Execute(frame *runtime.Frame) {
	frame.OperateStack().PushLong(1)
}

