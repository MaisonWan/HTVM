package runtime

import (
	"math"
)

// 同时容纳int和引用
type Slot struct {
	number int32
	ref    *Object
}

// 局部变量表
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

// 设置整形
func (self LocalVars) SetInt(index uint, value int32) {
	self[index].number = value
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].number
}

func (self LocalVars) SetFloat(index uint, value float32) {
	bits := math.Float32bits(value)
	self[index].number = int32(bits)
}

func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].number)
	return math.Float32frombits(bits)
}

// 设置长整形，需要两个int的空间
func (self LocalVars) SetLong(index uint, value int64) {
	self[index].number = int32(value)
	self[index + 1].number = int32(value >> 32)
}

// 得到长整型
func (self LocalVars) GetLong(index uint) int64 {
	low := self[index].number
	high := self[index + 1].number
	return (int64(high) << 32) | int64(low)
}

// 写入64位浮点
func (self LocalVars) SetDouble(index uint, value float64) {
	bits := math.Float64bits(value)
	self.SetLong(index, int64(bits))
}

// 得到浮点数
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

