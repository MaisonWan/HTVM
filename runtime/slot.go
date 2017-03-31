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
func (lv LocalVars) SetInt(index uint, value int32) {
	lv[index].number = value
}

func (lv LocalVars) GetInt(index uint) int32 {
	return lv[index].number
}

func (lv LocalVars) SetFloat(index uint, value float32) {
	bits := math.Float32bits(value)
	lv[index].number = int32(bits)
}

func (lv LocalVars) GetFloat(index uint) float32 {
	bits := uint32(lv[index].number)
	return math.Float32frombits(bits)
}

// 设置长整形，需要两个int的空间
func (lv LocalVars) SetLong(index uint, value int64) {
	lv[index].number = int32(value)
	lv[index + 1].number = int32(value >> 32)
}

// 得到长整型
func (lv LocalVars) GetLong(index uint) int64 {
	low := uint32(lv[index].number)
	high := uint32(lv[index + 1].number)
	return (int64(high) << 32) | int64(low)
}

// 写入64位浮点
func (lv LocalVars) SetDouble(index uint, value float64) {
	bits := math.Float64bits(value)
	lv.SetLong(index, int64(bits))
}

// 得到浮点数
func (lv LocalVars) GetDouble(index uint) float64 {
	bits := uint64(lv.GetLong(index))
	return math.Float64frombits(bits)
}

// 设置引用值
func (lv LocalVars) SetRef(index uint, obj *Object) {
	lv[index].ref = obj
}

// 获取引用值
func (lv LocalVars) GetRef(index uint) *Object {
	return lv[index].ref
}