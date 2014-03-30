// The sieve of Eratosthenes is one of the most efficient ways to find all of
// the smaller primes (below 10 million or so).

package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Println("Calculate number of primes less than or equal to what number? ")
  _, err := fmt.Scanf("%d", &n)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  if n < 1 {
    fmt.Println("Error:", n, "is too low - no primes exist less than 1")
    return
  } else if n > 100000000 {
      fmt.Println("Error:", n, "is too big!  Try a smaller number.")
      return
  }

  // initially assume all integers are prime
  // Take advantage of default bool value of false
  isNotPrime := make([]bool, n-1)

  // mark non-primes <= n using Sieve of Eratosthenes
  for i := 2; i*i <= n; i++ {

    // if i is prime, then mark multiples of i as nonprime
    // suffices to consider mutiples i, i+1, ..., N/i
    if !isNotPrime[i-2] {
      for j := i; i*j <= n; j++ {
        isNotPrime[i*j-2] = true;
      }
    }
  }

  // count primes
  primes := 0;
  for i := range(isNotPrime) {
    if !isNotPrime[i] {
      primes++;
    }
  }
  fmt.Println("The number of primes <=", n, "is", primes);
}
