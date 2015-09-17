package main

import (
	"testing"
)

// TestFibonnaci is an exported type that
// tests fibonaccical
func TestFibonnaci(t *testing.T) {
	result := calFib(6)
	if result != 8 {
		t.Errorf("The fibonacci of 6 is incorrect %d", result)
	}
	/*  elseif{
	    fmt.Println("Calculated Correctly")
	  }*/
}
