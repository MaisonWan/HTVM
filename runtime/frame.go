package runtime

// 帧
type Frame struct {
	next         *Frame
	localVars    LocalVars
	operateStack *OperateStack
	thread       *Thread
	nextPc       int
}

// 创建新的帧，最大的本地，最大的栈空间
func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:      thread,
		localVars:   newLocalVars(maxLocals),
		operateStack:newOperateStack(maxStack),
	}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) OperateStack() *OperateStack {
	return f.operateStack
}

func (f *Frame) CurrentThread() *Thread {
	return f.thread
}

func (f *Frame) NextPc() int {
	return f.nextPc
}

func (f *Frame) SetNextPc(nextPc int) {
	f.nextPc = nextPc
}