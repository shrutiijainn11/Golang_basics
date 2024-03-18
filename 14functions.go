package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 5, 8, 7, 6)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro message is: ", myMessage)
}

// func () {
// 	fmt.Println("Another method")
// }()  ---anonymous functions exist

func adder(valOne int, valTwo int) int { //value expected to be return goes here(datatype)
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Namastey from golang")
}
