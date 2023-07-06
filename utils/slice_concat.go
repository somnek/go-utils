package utils

func SliceConcat[T comparable](l []T, ls ...[]T) []T {
	result := append([]T{}, l...)
	for _, v := range ls {
		result = append(result, v...)
	}
	return result
}
