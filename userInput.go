// Created by: Mr Coxall
// Created on: Sep 2020
//
// This program does addition

package main

import "fmt"

func main() {
	// This function does addition
	var firstName string
	var age int

	// input
	fmt.Println("This program gets a user's name and age.")
	fmt.Println()
	fmt.Print("Enter your first name: ")
	fmt.Scanln(&firstName)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	// output
	fmt.Print("Your info is: " , firstName, ", age ", age, ".")
	fmt.Println("\n\nDone.")
}
