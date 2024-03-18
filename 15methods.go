package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang; no super or parent

	shruti := User{"Shruti", "shruti@go.dev", true, 20}
	fmt.Println(shruti)
	fmt.Println("Shruti details are: %+v\n", shruti)
	fmt.Println("Name is %v and email is%v.", shruti.Name, shruti.Email)
	shruti.GetStatus()
	shruti.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
