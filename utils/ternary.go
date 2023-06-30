package utils

// like js ternar:  ex ? a : b
func Ternary(ex bool, a, b string) string {
	if ex {
		return a
	}
	return b
}
