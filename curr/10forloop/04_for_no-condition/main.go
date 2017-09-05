package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
		/*no condition given to stop, this would run-on; press ^C to break out*/
	}
}
