// Find π to the Nth digit
// Since float64 can only hold ~16 decimal digits, we're going to need to use
// Go's math/big package for arbitrary precision arithmetic.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Printf("Compute π to how many decimal places? (≤200) ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if n > 200 {
		fmt.Println("Error:", n, "> 200 - too many decimal places!")
		return
	} else if n < 0 {
		fmt.Println("Error:", n, "< 0")
		return
	}

	/* Use Simon Plouffe's decimal digit extraction formula to calculate π.
	   For more information, see the following article:
	   Plouffe, Simon. "On the computation of the nth decimal digit of various
	   transcendental numbers." (1996) http://arxiv.org/pdf/0912.0303.pdf

	   The general formula used is shown below:
	                    ∞ n*2^n*(n!)^2
	           π + 3 =  Σ ------------
	                   n=1   (2*n)!
	*/
	pi := big.NewRat(0, 1)
	tempInt := big.NewInt(1)
	term := big.NewRat(0.0, 1.0)

	for k := 1; k < 4*n+10; k++ {
		top := big.NewInt(int64(k))
		tempInt = tempInt.Exp(big.NewInt(2), big.NewInt(int64(k)), nil)
		top = top.Mul(top, tempInt)
		tempInt = tempInt.Exp(tempInt.Set(fact(k)), big.NewInt(2), nil)
		top = top.Mul(top, tempInt)
		bottom := fact(2 * k)
		term := term.SetFrac(top, bottom)
		pi.Add(pi, term)
	}

	pi = pi.Add(pi, big.NewRat(-3, 1))
	fmt.Println("π to", n, "decimal places: ")
	fmt.Println(pi.FloatString(n))
}

// Simple factorial helper function.
// Takes an int but returns a *big.Int since factorial results are much much
// larger than their arguments.
func fact(a int) *big.Int {
	b := big.NewInt(1)

	for i := a; i > 1; i-- {
		b = b.Mul(b, big.NewInt(int64(i)))
	}

	return b
}
