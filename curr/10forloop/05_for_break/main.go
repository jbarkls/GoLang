package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		} /* NOTE: i >= 10 will print 1 ---> 10, because when i=9 it will drop to i++ making 10, and the Println is before the conditional check */
		i++
	}
}
