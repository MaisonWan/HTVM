package runtime

// jvm中栈指针，超出这个大小，就会出现stackoverflow
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame // 栈顶指针
}

// 创建新栈
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize:maxSize,
	}
}

// 入栈
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.next = self._top
	}
	self._top = frame
	self.size++
}

// 弹出栈空间
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	top := self._top
	self._top = top.next
	top.next = nil
	self.size--
	return top
}

// 返回栈顶空间，不弹出栈
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}