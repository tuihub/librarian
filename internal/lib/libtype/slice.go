package libtype

import (
	"math/rand"
)

func ShuffleSlice[T any](ss []T) {
	for i := range ss {
		j := rand.Intn(i + 1) //nolint:gosec // not critical
		ss[i], ss[j] = ss[j], ss[i]
	}
}

func DiffSlices[T comparable](a, b []T) ([]T, []T) {
	onlyA := make([]T, 0, len(a))
	onlyB := make([]T, 0, len(b))
	m := make(map[T]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := m[v]; !ok {
			onlyB = append(onlyB, v)
		}
	}
	for _, v := range a {
		if _, ok := m[v]; !ok {
			onlyA = append(onlyA, v)
		}
	}
	return onlyA, onlyB
}
