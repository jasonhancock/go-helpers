package helpers

type UniqueList[T comparable] struct {
	data []T
	seen map[T]struct{}
}

func NewUniqueList[T comparable]() *UniqueList[T] {
	return &UniqueList[T]{
		seen: make(map[T]struct{}),
	}
}

func (u *UniqueList[T]) Add(data ...T) {
	for _, element := range data {
		if _, ok := u.seen[element]; ok {
			continue
		}
		u.data = append(u.data, element)
		u.seen[element] = struct{}{}
	}
}

func (u *UniqueList[T]) Data() []T {
	return u.data
}
