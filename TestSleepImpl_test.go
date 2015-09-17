package main

import (
	"fmt"
	"testing"
	"time"
)

//TestSleep to test userSleep
func TestSleep(t *testing.T) {
	when := time.Now()
	userSleep(5)
	deadline := when.Add(5 * 1000000000 * time.Nanosecond)
	if time.Now().Nanosecond() > deadline.Nanosecond() {
		fmt.Println("sucess")
	}
	userSleep(10)
	deadline = when.Add(10 * 1000000000 * time.Nanosecond)
	if time.Now().Nanosecond() > deadline.Nanosecond() {
		fmt.Println("sucess")
	}
}
