package constants

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *runtime.Frame) {
	// 不用做
}