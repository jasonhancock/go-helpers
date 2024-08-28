package helpers

// Contains returns true if needle is present in haystack.
func Contains[T comparable](haystack []T, needle T) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}
