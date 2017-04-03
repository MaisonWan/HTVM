package instructions

import (
	_ "HTVM/instructions/base"
	"HTVM/instructions/constants"
	"HTVM/instructions/loads"
	"HTVM/instructions/stores"
	"HTVM/instructions/stack"
	"HTVM/instructions/math"
	"HTVM/instructions/conversions"
	"HTVM/instructions/comparisons"
	"HTVM/instructions/base"
	"fmt"
	"HTVM/instructions/control"
)

var (
	nop           = &constants.NOP{}
	aconst_null   = &constants.ACONST_NULL{}
	iconst_m1     = &constants.ICONST_M1{}
	iconst_0      = &constants.ICONST_0{}
	iconst_1      = &constants.ICONST_1{}
	iconst_2      = &constants.ICONST_2{}
	iconst_3      = &constants.ICONST_3{}
	iconst_4      = &constants.ICONST_4{}
	iconst_5      = &constants.ICONST_5{}
	lconst_0      = &constants.LCONST_0{}
	lconst_1      = &constants.LCONST_1{}
	fconst_0      = &constants.FCONST_0{}
	fconst_1      = &constants.FCONST_1{}
	fconst_2      = &constants.FCONST_2{}
	dconst_0      = &constants.DCONST_0{}
	dconst_1      = &constants.DCONST_1{}
	iload_0       = &load.ILOAD_0{}
	iload_1       = &load.ILOAD_1{}
	iload_2       = &load.ILOAD_2{}
	iload_3       = &load.ILOAD_3{}
	lload_0       = &load.LLOAD_0{}
	lload_1       = &load.LLOAD_1{}
	lload_2       = &load.LLOAD_2{}
	lload_3       = &load.LLOAD_3{}
	fload_0       = &load.FLOAD_0{}
	fload_1       = &load.FLOAD_1{}
	fload_2       = &load.FLOAD_2{}
	fload_3       = &load.FLOAD_3{}
	dload_0       = &load.DLOAD_0{}
	dload_1       = &load.DLOAD_1{}
	dload_2       = &load.DLOAD_2{}
	dload_3       = &load.DLOAD_3{}
	aload_0       = &load.ALOAD_0{}
	aload_1       = &load.ALOAD_1{}
	aload_2       = &load.ALOAD_2{}
	aload_3       = &load.ALOAD_3{}
	iaload        = &load.IALOAD{}
	laload        = &load.LALOAD{}
	faload        = &load.FALOAD{}
	daload        = &load.DALOAD{}
	aaload        = &load.AALOAD{}
	baload        = &load.BALOAD{}
	caload        = &load.CALOAD{}
	saload        = &load.SALOAD{}
	istore_0      = &stores.ISTORE_0{}
	istore_1      = &stores.ISTORE_1{}
	istore_2      = &stores.ISTORE_2{}
	istore_3      = &stores.ISTORE_3{}
	lstore_0      = &stores.LSTORE_0{}
	lstore_1      = &stores.LSTORE_1{}
	lstore_2      = &stores.LSTORE_2{}
	lstore_3      = &stores.LSTORE_3{}
	fstore_0      = &stores.FSTORE_0{}
	fstore_1      = &stores.FSTORE_1{}
	fstore_2      = &stores.FSTORE_2{}
	fstore_3      = &stores.FSTORE_3{}
	dstore_0      = &stores.DSTORE_0{}
	dstore_1      = &stores.DSTORE_1{}
	dstore_2      = &stores.DSTORE_2{}
	dstore_3      = &stores.DSTORE_3{}
	astore_0      = &stores.ASTORE_0{}
	astore_1      = &stores.ASTORE_1{}
	astore_2      = &stores.ASTORE_2{}
	astore_3      = &stores.ASTORE_3{}
	iastore       = &stores.IASTORE{}
	lastore       = &stores.LASTORE{}
	fastore       = &stores.FASTORE{}
	dastore       = &stores.DASTORE{}
	aastore       = &stores.AASTORE{}
	bastore       = &stores.BASTORE{}
	castore       = &stores.CASTORE{}
	sastore       = &stores.SASTORE{}
	pop           = &stack.POP{}
	pop2          = &stack.POP2{}
	dup           = &stack.DUP{}
	dup_x1        = &stack.DUP_X1{}
	dup_x2        = &stack.DUP_X2{}
	dup2          = &stack.DUP2{}
	dup2_x1       = &stack.DUP2_X1{}
	dup2_x2       = &stack.DUP2_X2{}
	swap          = &stack.SWAP{}
	iadd          = &math.IADD{}
	ladd          = &math.LADD{}
	fadd          = &math.FADD{}
	dadd          = &math.DADD{}
	isub          = &math.ISUB{}
	lsub          = &math.LSUB{}
	fsub          = &math.FSUB{}
	dsub          = &math.DSUB{}
	imul          = &math.IMUL{}
	lmul          = &math.LMUL{}
	fmul          = &math.FMUL{}
	dmul          = &math.DMUL{}
	idiv          = &math.IDIV{}
	ldiv          = &math.LDIV{}
	fdiv          = &math.FDIV{}
	ddiv          = &math.DDIV{}
	irem          = &math.IREM{}
	lrem          = &math.LREM{}
	frem          = &math.FREM{}
	drem          = &math.DREM{}
	ineg          = &math.INEG{}
	lneg          = &math.LNEG{}
	fneg          = &math.FNEG{}
	dneg          = &math.DNEG{}
	ishl          = &math.ISHL{}
	lshl          = &math.LSHL{}
	ishr          = &math.ISHR{}
	lshr          = &math.LSHR{}
	iushr         = &math.IUSHR{}
	lushr         = &math.LUSHR{}
	iand          = &math.IAND{}
	land          = &math.LAND{}
	ior           = &math.IOR{}
	lor           = &math.LOR{}
	ixor          = &math.IXOR{}
	lxor          = &math.LXOR{}
	iinc          = &math.IINC{}
	i2l           = &conversions.I2L{}
	i2f           = &conversions.I2F{}
	i2d           = &conversions.I2D{}
	l2i           = &conversions.L2I{}
	l2f           = &conversions.L2F{}
	l2d           = &conversions.L2D{}
	f2i           = &conversions.F2I{}
	f2l           = &conversions.F2L{}
	f2d           = &conversions.F2D{}
	d2i           = &conversions.D2I{}
	d2l           = &conversions.D2L{}
	d2f           = &conversions.D2F{}
	i2b           = &conversions.I2B{}
	i2c           = &conversions.I2C{}
	i2s           = &conversions.I2S{}
	lcmp          = &comparisons.LCMP{}
	fcmpl         = &comparisons.FCMPL{}
	fcmpg         = &comparisons.FCMPG{}
	dcmpl         = &comparisons.DCMPL{}
	dcmpg         = &comparisons.DCMPG{}
	if_icmpgt     = &comparisons.IF_ICMP_GT{}
	go_to         = &control.GOTO{}
	bipush        = &constants.BIPUSH{}
//	ireturn       = &IRETURN{}
//	lreturn       = &LRETURN{}
//	freturn       = &FRETURN{}
//	dreturn       = &DRETURN{}
//	areturn       = &ARETURN{}
//	_return       = &RETURN{}
//	arraylength   = &ARRAY_LENGTH{}
//	athrow        = &ATHROW{}
//	monitorenter  = &MONITOR_ENTER{}
//	monitorexit   = &MONITOR_EXIT{}
//	invoke_native = &INVOKE_NATIVE{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00: return nop
	case 0x01: return aconst_null
	case 0x02: return iconst_m1
	case 0x03: return iconst_0
	case 0x04: return iconst_1
	case 0x05: return iconst_2
	case 0x06: return iconst_3
	case 0x07: return iconst_4
	case 0x08: return iconst_5
	case 0x09: return lconst_0
	case 0x0a: return lconst_1
	case 0x10: return bipush
	case 0x1a: return iload_0
	case 0x1b: return iload_1
	case 0x1c: return iload_2

	case 0x3b: return istore_0
	case 0x3c: return istore_1
	case 0x3d: return istore_2

	case 0x60: return iadd
	case 0x84: return iinc

	case 0xa3: return if_icmpgt
	case 0xa7: return go_to

	//case 0xb2: return getstati
	default:
		panic(fmt.Errorf("Unsupported opcde: 0x%x !", opcode))
	}
}