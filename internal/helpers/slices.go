package helpers

func Convert[T, N any](slice []T, f func(v T) N) []N {
	res := make([]N, 0, len(slice))
	for _, v := range slice {
		res = append(res, f(v))
	}
	return res
}
