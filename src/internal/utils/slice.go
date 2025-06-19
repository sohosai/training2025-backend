package utils

func Map[T interface{}, U interface{}](slice []U, op func(U) T) []T {
	result := []T{}
	for _, v := range slice {
		result = append(result, op(v))
	}

	return result
}
