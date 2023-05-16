// Created By: Charlotte Jhu
// Created on: May 2023
//
// This program does division

package main

import (
	"fmt"
)

func main() {
	// This function does division
	// variables
	var number1 int
	var number2 int
	var counter = 0

	// input
	fmt.Println("Let's do some division!")
	fmt.Println()
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&number1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&number2)
	fmt.Println()

	// process
	for number1 >= number2 {
		number1 -= number2
		counter++
	}

	// output
	if number1 != 0 {
		fmt.Println("The answer is", counter, "with a remainder of", number1)
	} else {
		fmt.Println("The answer is", counter)
	}
	fmt.Println("\nDone.")
}
