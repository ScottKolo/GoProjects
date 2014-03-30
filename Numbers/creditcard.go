package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the credit card number to be validated: ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ccNum := make([]int, 19, 19)
	i := 0
	for _, rune := range input {
		n, err := strconv.Atoi(string(rune))
		if err == nil {
			ccNum[i] = n
			i++
		}
	}

	ccNum = ccNum[:i]

	valid, cardType := ValidateCreditCard(ccNum)
	if valid && cardType != "Unknown" {
		fmt.Printf("Credit card number is a valid %s card\n", cardType)
	} else {
		fmt.Println("Credit card number is NOT valid!")
	}
}

func ValidateCreditCard(ccNum []int) (valid bool, cardType string) {
	valid = luhnCheck(ccNum)
	cardType = getCardType(ccNum)
	return valid, cardType
}

func getCardType(ccNum []int) string {
	var sArray = make([]string, len(ccNum))
	for i, n := range ccNum {
		sArray[i] = strconv.Itoa(n)
	}
	s := strings.Join(sArray, "")
	match, _ := regexp.MatchString("^3[47][0-9]{13}$", s)
	if match {
		return "American Express"
	}
	match, _ = regexp.MatchString("^4[0-9]{12}(?:[0-9]{3})?$", s)
	if match {
		return "Visa"
	}
	match, _ = regexp.MatchString("^5[1-5][0-9]{14}$", s)
	if match {
		return "MasterCard"
	}
	match, _ = regexp.MatchString("^3(?:0[0-5]|[68][0-9])[0-9]{11}$", s)
	if match {
		return "Diners Club"
	}
	match, _ = regexp.MatchString("^6(?:011|5[0-9]{2})[0-9]{12}$", s)
	if match {
		return "Discover"
	}
	match, _ = regexp.MatchString("^(?:2131|1800|35[0-9]{3})[0-9]{11}$", s)
	if match {
		return "JCB"
	}
	return "Unknown"
}

func luhnCheck(ccNum []int) bool {
	checksum := 0
	for i, n := range ccNum {
		if (i+len(ccNum)%2)%2 == 0 {
			checksum += sumDigits(2 * n)
		} else {
			checksum += n
		}
	}

	return checksum%10 == 0
}

func sumDigits(n int) int {
	hundreds := n / 100
	tens := (n - 100*hundreds) / 10
	ones := n - 100*hundreds - 10*tens
	return ones + tens + hundreds
}
