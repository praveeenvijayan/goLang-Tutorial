package main

import "fmt"

func main() {
	// defer prints in lifo

	// output of below code
	// Defer in go
	// Heyyyyyyyyy
	// Hello
	// Hello world
	defer fmt.Println("Hello world")
	defer fmt.Println("Hello")
	fmt.Println("Defer in go")
	defer fmt.Println("Heyyyyyyyyy")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}	
