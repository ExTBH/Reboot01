package execute

import "slices"

func push[T any](from, to []T) ([]T, []T) {
	return from[1:], slices.Insert(to, 0, from[0])
}
