package diffsquares

func SquareOfSum(n int) int {
	i := (1 + n) * n / 2

	return i * i
}

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
