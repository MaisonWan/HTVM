package constants

import (
	"HTVM/instructions/base"
	"HTVM/runtime"
)

type Nop struct {
	base.NoOperandsInstruction
}

func (self *Nop) Execute(frame *runtime.Frame) {
	// 不用做
}