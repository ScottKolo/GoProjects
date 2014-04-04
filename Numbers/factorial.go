// The Factorial of a positive integer, n, is defined as the product of the
// sequence n, n-1, n-2, ...1 and the factorial of zero, 0, is defined as
// being 1. Solve this using both loops and recursion.

package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Compute factorial for what number?")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if n > 25 {
		fmt.Println("Error:", n, "is too big! Try an integer in [1, 25].")
		return
	}
	if n < 0 {
		fmt.Println("Error:", n, "is too small! Try an integer in [1, 25].")
		return
	}

	fmt.Printf("Loop implementation:      %d! = %d\n", n, factorialLoop(n))
	fmt.Printf("Recursive implementation: %d! = %d\n", n, factorialRecursive(n))
}

func factorialLoop(n int) int {
	if n == 0 {
		return 1
	}
	result := n
	for i := n - 1; i > 1; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}
