package utils

func Keys[K comparable, V any](data map[K]V) []K {
	result := make([]K, 0, len(data))
	for k := range data {
		result = append(result, k)
	}

	return result
}
