// Have the user enter a number and find all Prime Factors (if there are any)
// and display them.

package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Compute prime factors of what number?")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if n < 2 {
		fmt.Println("Error:", n, "is too small! Try an integer n ≥ 2.")
		return
	}

	fmt.Printf("Prime factorization of %d:\n", n)
	fmt.Printf("%d = ", n)
	first := true
	for i := 2; i <= n; i++ {
		// Looks like it's prime - save some time and skip to the end
		if i > n/2 && first {
			fmt.Printf("%d (number is prime!)", n)
			i = n
		} else if n%i == 0 {
			if first {
				fmt.Print(i)
				first = false
			} else {
				fmt.Printf(" * %d", i)
			}
			n /= i
			i = 1
		}

	}
	fmt.Println()
}
