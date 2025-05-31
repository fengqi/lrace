package lrace

// Ternary 三元运算
func Ternary[T any](expr bool, left, right T) T {
	if expr {
		return left
	}
	return right
}
