package helpers

func Ternary[T any](cond bool, trueCond T, falseCond T) T {
	if cond {
		return trueCond
	}
	return falseCond
}

func ValueOrDefault[T comparable](val, defaultVal T) T {
	var zeroVal = new(T)
	if val == *zeroVal {
		return defaultVal
	}
	return val
}
