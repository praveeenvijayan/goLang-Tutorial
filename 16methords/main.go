package main

import "fmt"

func main() {
	fmt.Println("Structs to go lang")
	// no inheritance in go lang
	praveen := User{"Praveen", "Praveen@lco.dev", true, 22}
	fmt.Println(praveen)
	fmt.Printf(" Praveen details are: %v\n ", praveen)
	fmt.Printf("Praveen details are: %+v\n", praveen)
	fmt.Printf("hello %v and your new email is %v.\n", praveen.Name, praveen.Email)
	praveen.GetStatus()
	praveen.NewEmail()
	fmt.Printf("hello %v and your new email is %v.\n", praveen.Name, praveen.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of user is ", u.Email)

}
