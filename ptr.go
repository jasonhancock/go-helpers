package helpers

// Ptr returns a pointer to any object.
func Ptr[T any](v T) *T {
	return &v
}
