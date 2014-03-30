// Have the program find prime numbers until the user chooses to stop asking
// for the next one.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input := "\n"

	// initially assume all integers are prime
	// Take advantage of default bool value of false
	// Hope the user won't hold down enter for more than half an hour
	n := 100000
	isNotPrime := make([]bool, n-1)

	// mark non-primes <= n using Sieve of Eratosthenes
	for i := 2; i*i <= n; i++ {

		// if i is prime, then mark multiples of i as nonprime
		// suffices to consider mutiples i, i+1, ..., N/i
		if !isNotPrime[i-2] {
			for j := i; i*j <= n; j++ {
				isNotPrime[i*j-2] = true
			}
		}
	}

	i := 2
	fmt.Println("Press enter for next prime - enter STOP to end")
	for input == "\n" {
		for isNotPrime[i-2] {
			i++
		}
		fmt.Printf("%d\n", i)
		i++
		input, _ = reader.ReadString('\n')
	}
}
