package native

// #include "float.h"
import "C"

var szfloat int

func SzFloat() int {
	return szfloat
}

func init() {
	szfloat = int(C.SzFloat())
}
