// Calculate the monthly payments of a fixed term mortgage over given Nth terms
// at a given interest rate.

package main

import (
	"fmt"
	"math"
)

func main() {
	var apr float64 // Mortgage rate
	var P float64   // Principal borrowed
	var Y float64   // Years

	fmt.Println("Enter the principal owed ($): ")
	_, err := fmt.Scanf("%f", &P)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Enter the mortgage rate or APR (%): ")
	_, err = fmt.Scanf("%f", &apr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Enter the length of the mortgage (years): ")
	_, err = fmt.Scanf("%f", &Y)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	n := math.Ceil(Y * 12) // Calculate payments from years
	// assuming monthly payments
	r := apr / 12
	A := P * (r / 100 * math.Pow(1+r/100, n)) / (math.Pow(1+r/100, n) - 1)

	fmt.Printf("Monthly mortgage payment for $%.2f at an APR of ", P)
	fmt.Printf("%.2f%% over %.0f years:\n", apr, Y)
	fmt.Printf("$%.2f \n", A)
}
