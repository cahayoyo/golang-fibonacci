package FibonacciPackage

func MemoizationFibonacci(n int, F map[int]int) int {
	if n <= 1 {
		F[n] = n
		return n
	}

	if _, ok := F[n-1]; !ok {
		F[n-1] = MemoizationFibonacci(n-1, F)
	}
	if _, ok := F[n-2]; !ok {
		F[n-2] = MemoizationFibonacci(n-2, F)
	}
	return F[n-1] + F[n-2]
}

func Memoization(n int) int {
	F := make(map[int]int)
	bucket := make([]int, n)

	for i := 1; i <= n; i++ {
		bucket[i-1] = MemoizationFibonacci(i, F)
	}

	result := bucket[n-1]
	return result
}
