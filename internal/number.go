package internal

// Number is a generic type for Go/C types that can be used as a number.
// It could be anything that you can convert to that type (e.g. C.int is a Number,
// because it can be directly converted to int)
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~bool // in C bool is technically a number
}

// WrapNumberPtr is a generic method to convert GOTYPE (int32/float32 e.t.c.) into CTYPE (c_int/c_float e.t.c.)
func WrapNumberPtr[CTYPE Number, GOTYPE Number](goValue *GOTYPE) (wrapped *CTYPE, finisher func()) {
	return ReinterpretCast[*CTYPE](goValue), func() {}
}
