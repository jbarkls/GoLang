package main

import "fmt"

func main() {

	a := 43

	fmt.Println(" [init 'a' to 43: 'a := 43', show value & addr..]")  // 43
	fmt.Println("1: val of 'a'=",a)  // 43
	fmt.Println("2: addr of 'a'= ",&a) // 0x20818a220

	var b = &a
	//fmt.Println()
	fmt.Println("\n ['b' becomes a pointer to addr of 'a' by 'var b = &a']")  // 0x20818a220
	fmt.Println("3: val of 'b' is 'addr of a', so b =",b)  // 0x20818a220
	fmt.Printf("4: addr of 'b' = %v",&b)  // 0x20818a220
	fmt.Println("  {note diff from val of 'b' above}")  // 0x20818a220
	fmt.Println("5: value at addr of 'b'[via '*b']=",*b) // 43
	fmt.Printf("  [note: '*' is 'dereferencing', and, an operator in '*b'\n   so, by use of *b we are pulling value from the\n   referenced addr (ie: val a) not -b's- value of addr of 'a']")
	fmt.Printf("\n\n")
	*b = 22        // b says, "The value at this address, change it to 42"
	fmt.Println(" [using '*b = 22' change the value at the referenced addr of 'b'\n which is stored at addr of 'a']") // 42
	fmt.Println("6: print 'a', now value of 'a'= ",a) // 42
	fmt.Println("  {compair to orig value at ln 1 & 5 above}") // 42

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value
}
