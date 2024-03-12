package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our Pizza: ")

	//comma ok syntax || error ok syntax

	input, _ := reader.ReadString('\n') //_ = when we don't care about it || error (if any error is encountered)
	fmt.Println("Thanks for rating, ", input)
	fmt.Println("Type of the rating is %T", input)

}
