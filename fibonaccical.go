package main

import "fmt"

func calFib(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	return calFib(x-1) + calFib(x-2)

}

func main() {

	var input int
	fmt.Println("Enter an integer number to calculate fibonacci: ")
	fmt.Scanf("%d", &input)

	if input < 0 {
		fmt.Print("Invalid Input")
		return
	}

	fmt.Println(calFib(input))

}
