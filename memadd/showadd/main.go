package main

import "fmt"

func main() {

	a := 43

	fmt.Println(" value of 'a' = ", a)
	fmt.Println(" Hex  address of 'a' = \n ", &a)
	fmt.Println(" dec  address of 'a' = ")
	fmt.Printf("  %d", &a)
}
