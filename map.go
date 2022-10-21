package helpers

import "errors"

// ErrKeyNotFound indicates a key was not present in the map.
var ErrKeyNotFound = errors.New("key not found")

// ErrIncorrectType indicates the item couldn't be asserted to the desired type.
var ErrIncorrectType = errors.New("element wasn't correct type")

// FromMap looks up a key from a map and does the type asserting necessary to
// return it as a typed value.
func FromMap[T any, K comparable](data map[K]any, key K) (T, error) {
	untyped, ok := data[key]
	if !ok {
		return *new(T), ErrKeyNotFound
	}
	str, ok := untyped.(T)
	if !ok {
		return *new(T), ErrIncorrectType
	}

	return str, nil
}
