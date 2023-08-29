package utils

func SliceFind[T comparable](l []T, x T) int {
	for i, x := range l {
		if x == x {
			return i
		}
	}
	return -1
}
