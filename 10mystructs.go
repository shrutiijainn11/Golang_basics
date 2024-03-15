package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang; no super or parent

	shruti := User{"Shruti", "shruti@go.dev", true, 20}
	fmt.Println(shruti)
	fmt.Println("Shruti details are: %+v\n", shruti)
	fmt.Println("Name is %v and email is%v.", shruti.Name, shruti.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
