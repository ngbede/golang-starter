package main // every go prgram must start with this line
// packages are go's way of organising and reusing our code
// 2 types of go programs are:
// executable: run directly from the command line
// libraries: (collection of code package together to be used in other programs).
// go is statically typed
import "fmt" // impport key inludes code from other packages to be used in a program

// functions are the building blocks of go code. func is the keyword
func main() { // main gets called when we run our programs
	fmt.Println("Hello world")
	fmt.Println("My name is Emmanuel")
}
