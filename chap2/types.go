package main

import "fmt"

func main() {
	fmt.Println("2 + 17 = ", 2+17)
	fmt.Println("22 / 7. = ", 22/7.)
	fmt.Println("This is a string via double quotes")
	fmt.Println(`This is also a string representition using backtiks`)
	fmt.Println(len("Emmanuel"))                  // print length of the string
	fmt.Println("EmmaNuel"[1])                    // print the 5th character in EmmaNuel i.e N
	fmt.Println("Emmanuel" + " Ngbede" + ` Sule`) // concatenation
	fmt.Println("2 * 2 = ", 2*2)
	fmt.Println("2 + 2 = ", 2+2)
	fmt.Println("3 - 2 = ", 3-2)
	fmt.Println("4 / 2 = ", 4/2)
	fmt.Println("9 % 4 = ", 9%4)
	// booleans:
	// &&: and symbol
	// ||: or symbol
	// !: not symbol
	// fmt.Println(true && true)
	// fmt.Println(false && true)
	// fmt.Println(false && false)
	// fmt.Println(true || false)
	// fmt.Println(false || false)
	// fmt.Println(!true)
}
