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

	fmt.Print("Enter the first integer to add: ")
	fmt.Scanln(&firstInteger)
	fmt.Print("Enter the second integer to add: ")
	fmt.Scanln(&secondInteger)
	fmt.Println(firstInteger, " + ", secondInteger, " = ", (firstInteger + secondInteger))
	fmt.Println("\nDone.")
}
