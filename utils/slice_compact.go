package utils

func SliceCompact[T comparable](l []T) []T {
	var zero T
	result := []T{}
	for _, v := range l {
		if v != zero {
			result = append(result, v)
		}
	}
	return result
}
