package main

import "fmt"

func main() {
	fmt.Println("Structs to go lang")
	// no inheritance in go lang

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
