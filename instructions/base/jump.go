package base

import "HTVM/runtime"

// 根据偏移量跳转
func Jump(frame *runtime.Frame, offset int) {
	pc := frame.CurrentThread().PC()
	nextPc := pc + offset
	frame.SetNextPc(nextPc)
}
