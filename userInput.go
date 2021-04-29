// Created by: Mr Coxall
// Created on: Sep 2020
//
// This program does addition

package main

import "fmt"

func main() {
	// This function does addition
	var firstInteger int
	var secondInteger int
	var sum int

	// input
	fmt.Println("This program adds 2 numbers together.")
	fmt.Println()
	fmt.Print("Enter the first integer to add: ")
	fmt.Scanln(&firstInteger)
	fmt.Print("Enter the second integer to add: ")
	fmt.Scanln(&secondInteger)

	// process
	sum = firstInteger + secondInteger

	// output
	fmt.Println(firstInteger, " + ", secondInteger, " = ", sum)
	fmt.Println("\nDone.")
}
