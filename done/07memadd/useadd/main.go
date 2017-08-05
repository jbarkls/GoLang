package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)	/* Here we scan (poll) for user input, and push it into the var meters vis &meters, its addr. */
	yards := meters * metersToYards
	fmt.Println(meters, " \n meters is ", yards, " yards.")
}
