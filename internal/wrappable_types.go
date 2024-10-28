package internal

import "unsafe"

// WrappableType represents a GO type that can be converted into a C value
// and back - into a GO value.
// CTYPE represents the equivalent C type of self.
// self is the type WrappableType applies to - TODO - figure out if it can be ommited :-)
// intentional values:
// - CTYPE is e.g. C.ImVec2, C.ImColor e.t.c.
// - self is a pointer type (e.g. *Vec2, Color)
type WrappableType[CTYPE any, self any] interface {
	// ToC converts self into CTYPE
	ToC() CTYPE
	// FromC force-converts taken any to CTYPE, converts it into self,
	// applies to receiver and returns it.
	FromC(unsafe.Pointer) self
}

// Wrap takes a variable of one of the types defined in this file
// and returns a pointer to its C equivalend as well as a "finisher" func.
// This finisher func should be called after doing any operations
// on C pointer in order to apply the changes back to the original GO variable.
func Wrap[CTYPE any, self any](in WrappableType[CTYPE, self]) (cPtr *CTYPE, finisher func()) {
	if in != nil {
		c := in.ToC()
		cPtr = &c

		finisher = func() {
			in.FromC(unsafe.Pointer(cPtr))
		}
	} else {
		finisher = func() {}
	}
	return cPtr, finisher
}
