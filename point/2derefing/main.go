package main

import "fmt"

func main() {

	a := 43

	fmt.Println("val of 'a' =",a)  // 43
	fmt.Println("addr of 'a' =",&a) // 0x20818a220

	var b = &a
	/* the above code makes b a pointer to the memory address where an int is stored b is of type "int pointer" *int -- the * is part of the type -- b is of type *int*/
	fmt.Printf("  With 'var b = &a'; 'b' a pointer of type int (*int)\n  to addr of 'a': \n")
	fmt.Printf("  thus val of 'b' = %v\n\n",b)  // addr of 'a'
	fmt.Printf("  now dereference 'b' [Println(*b)] to get\n  val @ addr ref'd by 'b' which is the value of 'a': \n") // val of 'a'
	fmt.Printf("  so *b = %v: ...the orig val of 'a'",*b)


	/* b is an int pointer;
	   b points to the memory address where an int is stored
	   to see the value in that memory address, add a * in front of b
	   this is known as de-referencing
	   the * is an operator in this case*/
}
