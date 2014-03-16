// Show how to spell out a number in English. You can use a preexisting
// implementation or roll your own, but you should support inputs up to at
// least one million.

package main

import "fmt"

func main() {
	fmt.Println("Enter an integer (≤1,000,000): ")
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err, "\n")
		return
	}
	if n > 1000000 || n < -1000000 {
		fmt.Printf("Error:|%d| > 1000000\n", n)
		return
	}

	fmt.Println(n, "spelled out in English:")
	fmt.Println(GetNumberName(n))
}

func GetNumberName(n int) string {

	// Keep track if it's a negative - we'll tack on a "negative" later
	neg := false

	if n < 0 {
		neg = true
		n = -n
	}

	name := ""

	if exceptionsMap[n] != "" {
		// Check right away if it's an exception
		name = exceptionsMap[n]
	} else if n > 0 && n < 10 {
		name = onesMap[n]
	} else if n >= 20 && n < 100 {
		tens := n / 10
		ones := n - tens*10
		name = tensMap[tens*10]
		if ones > 0 {
			name += "-" + onesMap[ones]
		}
	} else if n >= 100 && n < 1000 {
		hundreds := n / 100
		tensAndLower := n - hundreds*100
		name = onesMap[hundreds] + " hundred"
		if tensAndLower > 0 {
			name += " " + GetNumberName(tensAndLower)
		}
	} else if n >= 1000 && n < 1000000 {
		thousands := n / 1000
		hundredsAndLower := n - thousands*1000
		name = GetNumberName(thousands) + " thousand"
		if hundredsAndLower >= 100 {
			name += ", " + GetNumberName(hundredsAndLower)
		} else if hundredsAndLower > 0 {
			name += " " + GetNumberName(hundredsAndLower)
		}
	}

	if neg {
		name = "negative " + name
	}

	return name
}

// Map for exceptions that don't follow the rules above
var exceptionsMap = map[int]string{
	0:       "zero",
	10:      "ten",
	11:      "eleven",
	12:      "twelve",
	13:      "thirteen",
	14:      "fourteen",
	15:      "fifteen",
	16:      "sixteen",
	17:      "seventeen",
	18:      "eighteen",
	19:      "nineteen",
	1000000: "one million",
}

var tensMap = map[int]string{
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

var onesMap = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}
