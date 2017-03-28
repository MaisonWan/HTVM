package runtime

// 帧
type Frame struct {
	next         *Frame
	localVars    LocalVars
	operateStack *OperateStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:newLocalVars(maxLocals),
		operateStack:newOperateStack(maxStack),
	}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) OperateStack() *OperateStack {
	return f.operateStack
}
