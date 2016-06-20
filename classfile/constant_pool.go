package classfile

// 常量池
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := reader.readUint16()
	cp := make([]ConstantInfo, cpCount)

	// 索引从1开始
	for i := 1; i < cpCount; i++ {

	}
}
