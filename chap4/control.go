// control flow in go
package main

import "fmt"

func main() {
	// For loops
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	} // OR
	fmt.Println("Second loop below")
	for i := 11; i < 21; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//if, else if and else control structure
	var (
		// //bmi1 = 17.412 underweight
		// bmi2 = 24.901 //normal weight
		bmi3 = 29.503 // overweight
	)
	myBmi := bmi3

	if myBmi < 18.0 {
		fmt.Println("You are underweight eat some food")
	} else if (myBmi > 18.0) && (myBmi <= 26.0) {
		fmt.Println("You have a normal weight keep going go the gym")
	} else {
		fmt.Println("You are overweight, you might need a bariatric surgery")
	}

	// switch cases
	switch 3 {
	case 0:
		fmt.Println("+")
	case 1:
		fmt.Println(" ")
	case 2:
		fmt.Println("ABC")
	case 3:
		fmt.Println("DEF")
	case 4:
		fmt.Println("GHI")
	case 5:
		fmt.Println("JKL")
	case 6:
		fmt.Println("MNO")
	case 7:
		fmt.Println("PQRS")
	case 8:
		fmt.Println("TUV")
	case 9:
		fmt.Println("WXYZ")
	default:
		fmt.Println("Invalid phone dial")
	}
	fizzbuzz()
}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 != 0) {
			fmt.Println("FIZZ")
		} else if (i%5 == 0) && (i%3 != 0) {
			fmt.Println("BUZZ")
		} else if (i%5 == 0) && (i%3 == 0) {
			fmt.Println("FIZZBUZZ")
		} else {
			fmt.Println(i)
		}
	}
}
