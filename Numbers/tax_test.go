// Attempted to create a test for the tax calculator, but Go package structure
// won't allow it.  Oh well!

package main_test

import (
	"testing"
)

// Tests for correctness
type testSet struct {
	price      float64
	taxRate    float64
	tax        float64
	totalPrice float64
}

var tests = []testSet{
	{100.0, 0.0, 0.0, 100.0},
	{0, 0.5, 0.0, 0.0},
	{100.0, 0.5, 50.0, 150.0},
}

func TestTaxCalculation(t *testing.T) {
	for _, testSet := range tests {
		totalPrice, tax := CalculateTax(testSet.price, testSet.taxRate)
		if totalPrice != testSet.totalPrice || tax != testSet.tax {
			t.Error(
				"For price", testSet.price,
				"and tax rate", testSet.taxRate,
				"expected tax of", testSet.tax,
				"and total price", testSet.totalPrice,
				"but got tax of", tax,
				"and total price", totalPrice,
			)
		}
	}
}
