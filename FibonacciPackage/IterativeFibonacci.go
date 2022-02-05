package FibonacciPackage

import "fmt"

func IterativeFibonacci(n int) int {
	t0 := 0
	t1 := 1
	var s, i int

	if n <= 1 {
		return n
	}
	for i = 2; i <= n; i++ {
		fmt.Println("In Loop", i, ",", "t0 variable is :", t0)
		fmt.Println("In Loop", i, ",", "t1 variable is :", t1)
		s = t0 + t1
		t0 = t1
		t1 = s
		fmt.Println("Sum of t0 + t1 :", s)
		fmt.Println()
	}
	return s
}
