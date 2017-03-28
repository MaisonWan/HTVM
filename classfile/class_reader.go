package classfile

import "encoding/binary"

// 按照字节读取class里数据
type ClassReader struct {
	data []byte
}

func (cr *ClassReader) ReadUint8() uint8 {
	value := cr.data[0]
	cr.data = cr.data[1:]
	return value
}

func (cr *ClassReader) ReadUint16() uint16 {
	value := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return value
}

func (cr *ClassReader) ReadUint32() uint32 {
	value := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return value
}

func (cr *ClassReader) ReadUint64() uint64 {
	value := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return value
}

func (cr *ClassReader) ReadUint16s() []uint16 {
	n := cr.ReadUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = cr.ReadUint16()
	}
	return s
}

func (cr *ClassReader) ReadBytes(n uint32) []byte {
	s := cr.data[:n]
	cr.data = cr.data[n:]
	return s
}