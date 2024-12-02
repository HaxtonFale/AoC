package math

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}