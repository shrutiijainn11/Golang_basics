package main

import "fmt"

const LoginToken string = "jkdjh" //first letter capital = public variable or constant

func main() {
	var username string = "shruti"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T \n", isLoggedin)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var smallfloat float32 = 255.2153671253613
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	var smallfloat64 float64 = 255.2153671253613
	fmt.Println(smallfloat64)
	fmt.Printf("Variable is of type: %T \n", smallfloat64)

	//default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit way of declaring variables
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000 //inside the method (here, main) we can use ':=' operator but not outside
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
