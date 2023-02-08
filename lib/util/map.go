package util

func GetKeysOrValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := []K{}
	values := []V{}

	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}
