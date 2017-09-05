package main

import "fmt"

func main() {
	var numOne int
	var numTwo int
	fmt.Printf("\n\nThis will Divide intergers\ngive remainder,\nand proof of operation:\n-----------------------\n")
	fmt.Print("Please enter a divdend: ")
	fmt.Scan(&numOne)
	fmt.Print("Please enter a divisor: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "/", numTwo, "It's quotient is", numOne/numTwo)
	if numOne%numTwo != 0 {
		fmt.Println("it's remainder is", numOne%numTwo)
	} else {
		fmt.Println("it has no remainder")
	}
	fmt.Println("-----------")
	fmt.Println("proof", numOne/numTwo, "*", numTwo, "=", numOne/numTwo*numTwo, "(Plus remainder of", numOne%numTwo, ")", "=", numOne, "\n\n")
}
