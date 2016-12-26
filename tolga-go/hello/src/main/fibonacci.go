package main

import "fmt"

func main() {
	fmt.Println("Fibonacci numbers by recursive function")
	for i := 0; i < 8; i++ {
		fmt.Print(recursiveFibonacci(i))
		fmt.Print(" ")
	}
	fmt.Println("\nFibonacci numbers by non-recursive function\n")
	nonRecursiveFibonacci(8)
}

func recursiveFibonacci(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 || n == 1 {
		return n
	}
	return recursiveFibonacci(n - 1) + recursiveFibonacci(n - 2)
}

func nonRecursiveFibonacci(n int)  {
	firstNumber := 0
	secondNumber := 1
	temp := 0
	for i :=0; i<n;i++{
		temp = firstNumber
		fmt.Print(firstNumber)
		firstNumber = secondNumber
		secondNumber += temp
		fmt.Print(" ")
	}
}
