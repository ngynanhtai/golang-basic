package common

func filterAny[T any](arr []T, f func(T) bool) []T {
	for i, v := range arr {
		if f(v) {
			return arr[i:]
		}
	}
	return arr
}

func mapAny[T any](arr []T, f func(T) T) []T {
	for i, v := range arr {
		arr[i] = f(v)
	}
	return arr
}
