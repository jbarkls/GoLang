package main

import "fmt"

func main() {
	var fname string
	var lname string
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&fname)
	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lname)
	fmt.Println("Hello")
	fmt.Printf("your name is %v, %v ", lname, fname)

}
