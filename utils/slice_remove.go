package utils

func SliceRemove[T comparable](l []T, i int) []T {
	if i >= len(l) {
		return l
	}
	return append(l[:i], l[i+1:]...)
}
