package main

import "fmt"

func zero1(z *int) {
	fmt.Println("func zero1 'z' addr =\n ", &z)
	*z = 0
	fmt.Println("func zero1 'z' value =\n addr. of x = ", z)

}
func zero2(z *int) {
	fmt.Println("func zero2 'z' addr =\n ",&z)
	*z = 0
	fmt.Println("func zero2 'z' value =\n addr. of y = ",z)
}

func main() {
	x := 5
	y := 6
	fmt.Println("'func main' addr of 'x':\n",&x)
	fmt.Println("'func Main' value of 'x' locally:\n",x)
	zero1(&x)
	fmt.Println("'func main' addr post zero1(&x) of 'x':\n",&x)
	fmt.Println("'func main' val 'x' post 'zero1(&x)' call:\n",x)
	fmt.Printf("\n\n")
	fmt.Println("'func main' addr of 'y':\n",&y)
	fmt.Println("func Main' value of 'y' locally:\n",y)
	zero2(&y)
	fmt.Println("'func main' addr post zero2(&y) of'y':\n",&y)
	fmt.Println("'func main' val 'y' post 'zero2(&y)' call:\n",y)

}
