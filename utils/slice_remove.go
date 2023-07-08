package utils

func SliceRemove[T comparable](l []T, x T) []T {
	result := []T{}
	for _, v := range l {
		if v != x {
			result = append(result, v)
		}
	}
	return result
}
