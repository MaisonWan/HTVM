package base

import "HTVM/runtime"

// 指令
type Instruction interface {
	// 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// 执行指令逻辑
	Execute(frame *runtime.Frame)
}

// 没有操作数的指令
type NoOperandsInstruction struct {
	
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader)  {
	// 没有指令，不需要做任何事情
}

// 跳转指令
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader)  {
	self.Offset = int(reader.ReadInt16())
}

// 单字节操作数
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index = uint(reader.ReadUint8())
}

// 双字节操作数
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index = uint(reader.ReadUint16())
}
