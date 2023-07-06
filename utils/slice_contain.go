package utils

func SliceContain[T comparable](l []T, x T) bool {
	for _, e := range l {
		if e == x {
			return true
		}
	}
	return false
}
