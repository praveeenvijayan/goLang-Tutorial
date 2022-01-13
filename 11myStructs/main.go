package main

import "fmt"

func main() {
	fmt.Println("Structs to go lang")
	// no inheritance in go lang
	praveen := User{"Praveen", "Praveen@lco.dev", true, 22}
	fmt.Println(praveen)
	fmt.Printf(" Praveen details are: %v\n ", praveen)
	fmt.Printf("Praveen details are: %+v\n", praveen)
	fmt.Printf("hello %v and your new email is %v\n.", praveen.Name, praveen.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
