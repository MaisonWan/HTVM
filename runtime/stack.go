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
func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if s._top != nil {
		frame.next = s._top
	}
	s._top = frame
	s.size++
}

// 弹出栈空间
func (s *Stack) pop() *Frame {
	if s._top == nil {
		panic("jvm stack is empty")
	}
	top := s._top
	s._top = top.next
	top.next = nil
	s.size--
	return top
}

// 返回栈顶空间，不弹出栈
func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}
	return s._top
}