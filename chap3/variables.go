package main

import "fmt"

var name string = "Emmanuel"

func main() {

	var age int8 = 20
	var weight float32 = 70.45
	var alive bool = true
	fmt.Println("My name is ", name)
	fmt.Println("I am", age, " years old")
	fmt.Println("I weigh", weight, "kilograms")
	fmt.Println("it is", alive, "that ", name, "is alive")
	var num1 int8 = 12
	var num2 int8 = 12
	fmt.Println(num1, " == ", num2, "=", num1 != num2)
	// dynamic declartion of varibles
	// the go compiler will infer the types of both x & y variables
	var x = 50
	y := "I was declared dynamically"
	fmt.Println(x)
	fmt.Println(y)
	add(4, 5)
	// constants start with the keword const
	const pi float32 = 3.142857
	fmt.Println(pi)
	// defining multiple variables
	var (
		product          = "vitamin C"
		quantity int8    = 32
		price    float32 = 50.25
	)
	fmt.Println("we have", quantity, "of", product, "which cost", price, "naira each")

	// //simple calc
	// fmt.Println("Input both numbers to add: ")
	// var n1 int8
	// var n2 int8
	// fmt.Scanf("%d %d", &n1, &n2)
	// result := n1 + n2
	// fmt.Println(n1, "+", n2, "=", result)

	main2() // run temperature converter here
}

func add(x int8, y int8) {
	fmt.Println(x + y)
	fmt.Println(name)
}

func main2() {

	//(C = (F − 32) * 5/9).
	//convert fahrenheit to celcius
	var fahrenheit = 0.
	fmt.Println("Input temperature in Fahrenheit: ")
	fmt.Scanf("%d", &fahrenheit)
	var celcius = (fahrenheit - 32.) * (5 / 9.)
	fmt.Println(fahrenheit, "Fahrenheit equal to", celcius, "celcius")

}
