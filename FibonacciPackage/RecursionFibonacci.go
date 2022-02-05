package FibonacciPackage

func RecursionFibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return RecursionFibonacci(n-2) + RecursionFibonacci(n-1)
	}
}
