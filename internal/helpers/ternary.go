package helpers

func Ternary[T any](cond bool, trueCond T, falseCond T) T {
	if cond {
		return trueCond
	}
	return falseCond
}
