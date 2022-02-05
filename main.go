package main

import (
	"fmt"
	"golang_fibonnaci/FibonacciPackage"
)

func main() {
	var number int
	fmt.Print("Enter Fibonnaci Number To Find : ")
	fmt.Scan(&number)

	fmt.Println("Using Recursion")
	fmt.Println("Result : ", FibonacciPackage.RecursionFibonacci(number))

	fmt.Println("Using Iterative")
	fmt.Println("Result : ", FibonacciPackage.IterativeFibonacci(number))
}
