// The user enters a cost and then the amount of money given. The program will
// figure out the change and the number of quarters, dimes, nickels, pennies
// needed for the change.

package main

import (
	"fmt"
	"math"
)

func main() {
	var price float64
	var payment float64

	fmt.Println("What is the price of the purchase? ($)")
	_, err := fmt.Scanf("%f", &price)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("What amount is given? ($)")
	_, err = fmt.Scanf("%f", &payment)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if payment < price {
		fmt.Printf("Not enough! $%.2f is less than $%.2f.", payment, price)
		return
	}

	diff := payment - price
	if diff >= 1000 {
		fmt.Println("Overpayment - do you have any smaller bills?")
		return
	}
	if math.Abs(diff) < 0.01 {
		fmt.Println("No change due - paid with exact change")
		return
	}

	// Bills
	hundreds := math.Floor(diff / 100)
	leftover := diff - hundreds*100
	fifties := math.Floor(leftover / 50)
	leftover -= fifties * 50
	twenties := math.Floor(leftover / 20)
	leftover -= twenties * 20
	tens := math.Floor(leftover / 10)
	leftover -= tens * 10
	fives := math.Floor(leftover / 5)
	leftover -= fives * 5
	ones := math.Floor(leftover)
	leftover -= ones

	// Coins (less than 1 dollar)
	leftover *= 100
	quarters := math.Floor(leftover / 25)
	leftover -= quarters * 25
	dimes := math.Floor(leftover / 10)
	leftover -= dimes * 10
	nickels := math.Floor(leftover / 5)
	leftover -= nickels * 5
	pennies := math.Floor(leftover + 0.5)

	fmt.Println("Change due:")
	if hundreds > 0 {
		fmt.Println("$100 bills:", hundreds)
	}
	if fifties > 0 {
		fmt.Println("$50 bills: ", fifties)
	}
	if twenties > 0 {
		fmt.Println("$20 bills: ", twenties)
	}
	if tens > 0 {
		fmt.Println("$10 bills: ", tens)
	}
	if fives > 0 {
		fmt.Println("$5 bills:  ", fives)
	}
	if ones > 0 {
		fmt.Println("$1 bills:  ", ones)
	}
	if quarters > 0 {
		fmt.Println("Quarters:  ", quarters)
	}
	if dimes > 0 {
		fmt.Println("Dimes:     ", dimes)
	}
	if nickels > 0 {
		fmt.Println("Nickels:   ", nickels)
	}
	if pennies > 0 {
		fmt.Println("Pennies:   ", pennies)
	}
}
