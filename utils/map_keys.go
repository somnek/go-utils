package utils

func MapKeys(m map[string]struct{}) []string{
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
