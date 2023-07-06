package utils

func SliceCount[T comparable](l []T, x T) int {
	var count int
	for _, v := range l {
		if v == x {
			count++
		}
	}
	return count
}
