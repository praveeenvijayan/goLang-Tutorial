package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome user Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rateing? ")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks ", input)
	fmt.Printf("Type of this rateing is %T", input)
}