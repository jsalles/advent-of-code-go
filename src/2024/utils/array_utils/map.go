package arrutils

func Map[T, U any](input []T, f func(T) U) []U {
	result := make([]U, len(input))
	for i, element := range input {
		result[i] = f(element)
	}
	return result
}
