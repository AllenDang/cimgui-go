package utils

import (
	"unsafe"
)

// SliceToPtr takes a slice and returns its C-compatible pointer.
func SliceToPtr[T any](slice []T) *T {
	if len(slice) == 0 {
		return nil
	}
	return &slice[0]
}

// PtrToSlice in short is an equvalent to `unsafe.Slice(x, length)`.
// It allows to restore imgui data presented in form of a pointer to a Go slice.
// Reference issue: https://github.com/AllenDang/cimgui-go/issues/160#issuecomment-2768647950
// Example usage:
//
//	Consider you have sortSpecs := imgui.TableGetSortSpecs()
//	Now, because sortSpecs is of type *imgui.TableSortSpecs, you can't directly convert it to the slice (this is NOT supposed to be slice - this is a pointer to the GO type).
//	You need to extract raw C pointer which is sortSpecs.CData.
//	After calling PtrToSLice(imgui.TableSortSpecs.CData, int(sortSpecs.Count)), you will need to call NewTableSortSpecsFromC on every element to get something useful.
func PtrToSlice[T any](x *T, length int) []T {
	if x == nil || length <= 0 {
		return nil
	}
	return unsafe.Slice(x, length)
}
