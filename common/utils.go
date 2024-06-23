package common

func filterAny[T any](arr []T, f func(T) bool) []T {
	var newArr []T
	for _, v := range arr {
		if f(v) {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

func mapAny[T any](arr []T, f func(T) T) []T {
	for i, v := range arr {
		arr[i] = f(v)
	}
	return arr
}
