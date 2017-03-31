package runtime

import "math"

// 操作数栈
type OperateStack struct {
	size  uint
	slots []Slot
}

// 创建一个新的操作数栈
func newOperateStack(maxStack uint) *OperateStack {
	if maxStack > 0 {
		return &OperateStack{
			slots:make([]Slot, maxStack),
		}
	}
	return nil
}

func (self *OperateStack) PushInt(value int32) {
	self.slots[self.size].number = value
	self.size++
}

func (self *OperateStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].number
}

func (self *OperateStack) PushFloat(value float32) {
	bits := math.Float32bits(value)
	self.slots[self.size].number = int32(bits)
	self.size++
}

func (self *OperateStack) PopFloat() float32 {
	self.size--
	n := uint32(self.slots[self.size].number)
	return math.Float32frombits(n)
}

func (self *OperateStack) PushLong(value int64) {
	self.slots[self.size].number = int32(value)
	self.slots[self.size + 1].number = int32(value >> 32)
	self.size += 2
}

func (self *OperateStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].number)
	high := uint32(self.slots[self.size + 1].number)
	return (int64(high) << 32) | int64(low)
}

func (self *OperateStack) PushDouble(value float64) {
	bits := math.Float64bits(value)
	self.PushLong(int64(bits))
}

func (self *OperateStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

// 引用类型
func (self *OperateStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}

// 弹出引用，记得置位为nil，内存释放
func (self *OperateStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return ref
}

// 引用类型
func (self *OperateStack) PushSlot(slot *Slot) {
	if slot != nil {
		self.slots[self.size] = slot
		self.size++
	}
}

// 弹出引用，记得置位为nil，内存释放
func (self *OperateStack) PopSlot() Slot {
	if self.size > 0 {
		self.size--
		return self.slots[self.size]
	}
	return nil
}
