package runtime

// 线程
type Thread struct {
	pc    int
	stack *Stack
}

// 创建新线程
func NewThread() *Thread {
	return &Thread{
		stack:newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

// 入栈一帧
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

// 弹出一帧
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// 当前栈顶帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(self, maxLocals, maxStack)
}