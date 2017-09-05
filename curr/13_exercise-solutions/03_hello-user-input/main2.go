package main

import "fmt"

func main() {
	var fname string
	var lname string
	fmt.Print("Please enter your first and last name: ")
	fmt.Scan(&fname, &lname)
	/*fmt.Print("Please enter your last name: ")
	fmt.Scan(&lname) */
	fmt.Println("Hello")
	fmt.Printf("%v, %v : \n Hmm... Interesting", lname, fname)

}
