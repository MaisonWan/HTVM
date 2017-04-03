package main

import (
	"HTVM/classfile"
	"HTVM/runtime"
	"fmt"
	"HTVM/instructions/base"
	"HTVM/instructions"
)

func interpreter(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttibutes()
	maxLocals := uint(codeAttr.MaxLocals())
	maxStack := uint(codeAttr.MaxStack())
	bytecode := codeAttr.Code()
	
	thread := runtime.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	
	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *runtime.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("Local Vars:%v\n", frame.LocalVars())
		fmt.Printf("Operate Stack:%v\n", frame.OperateStack())
		panic(r)
	}
}

func loop(thread *runtime.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPc()
		thread.SetPC(pc)
		
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		//fmt.Printf("opcode = %v\n", opcode)
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPc(reader.PC())

		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}