package util

// Fact returns the factorial of n
func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}
