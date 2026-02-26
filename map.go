package helpers

import (
	"cmp"
	"errors"
	"slices"
)

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

// Keys returns a slice of keys from the provided map. Does not sort them, order
// is not deterministic.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// KeysSorted returns a sorted slice of keys from the provided map.
func KeysSorted[K cmp.Ordered, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
