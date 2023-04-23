package utils

func Prepend[T any](slice []T, item T) []T {
	slice = append(slice, item)
	copy(slice[1:], slice)
	slice[0] = item
	return slice
}
