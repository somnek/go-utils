package utils

func SliceFind[T comparable](l []T, v T) int {
	for i, x := range l {
		if x == v {
			return i
		}
	}
	return -1
}
