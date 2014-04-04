// This program will get the true atomic time from an atomic time clock on the
// Internet. Use any one of the atomic clocks returned by a simple Google search

package main

import (
	"fmt"
	"github.com/beevik/ntp" // Using pre-made NTP protocol implementation
)

func main() {
	time, err := ntp.Time("time.nist.gov")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	const layout = "3:04:05 PM (MST) on Monday, January _2, 2006"
	fmt.Println("Current Local Time:")
	fmt.Println(time.Local().Format(layout))
}
