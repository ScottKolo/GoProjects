// Tax Calculator - Asks the user to enter a cost and either a country or
// state tax. It then returns the tax plus the total cost with tax.

package main

import "fmt"

func main() {
	var p float64
	var t float64
	fmt.Println("What is the base price of the purchase? ($)")
	_, err := fmt.Scanf("%f", &p)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("What is the tax rate? (%)")
	_, err = fmt.Scanf("%f", &t)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("A purchase of $%.2f with a tax rate of %.2f%%\n", p, t)
	totalPrice, tax := CalculateTax(p, t/100)
	fmt.Printf("Total Tax: $%.2f\n", tax)
	fmt.Printf("Total Purchase Price: $%.2f\n", totalPrice)
}

func CalculateTax(price float64, taxRate float64) (totalPrice float64, tax float64) {
	tax = price * taxRate
	totalPrice = price + tax

	return totalPrice, tax
}
