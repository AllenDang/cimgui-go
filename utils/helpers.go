package utils

// SliceToPtr takes a slice and returns its C-compatible pointer.
func SliceToPtr[T any](slice []T) *T {
	if len(slice) == 0 {
		return nil
	}
	return &slice[0]
}
