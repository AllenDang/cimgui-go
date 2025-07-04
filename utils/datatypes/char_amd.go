//go:build amd64

package datatypes

// Char represents underlying type of C.char. On amd64 it is equivalent to int8, on arm64 it is uint8.
// See https://github.com/AllenDang/cimgui-go/issues/140
type Char int8
