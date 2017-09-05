package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 { // even returns to i++
			continue // odd drops to Println()
			/*if i%2 != 0
			continue  	.. != is opposite == to print evens */
		}
		fmt.Println(i) //printing only odds
		if i >= 50 {   // until cond = true -> i++
			break // when above true, end loop.
		}
	}
}

// i%[num] is give remainder of i/[num]
