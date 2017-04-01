package main

import (
	"HTVM/classfile"
	"HTVM/runtime"
	"fmt"
	"HTVM/instructions/base"
)

func interpreter(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttibutes()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	
	thread := runtime.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	
	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame runtime.Frame) {
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
	}
}