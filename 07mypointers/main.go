package main

import "fmt"

func main() {
	fmt.Println("welcome to class on pointers: ")

	var ptrs *int
	fmt.Println("value of Ponter is: ", ptrs)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("value of Pointer is: ", ptr)
	fmt.Println("value of Pointer is: ", *ptr)

	*ptr = *ptr * 2

	fmt.Println("My New Value is: ", myNumber)

}
