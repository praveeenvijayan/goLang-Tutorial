package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	var result string

	loginCount := 10

	if loginCount < 10 {
		result = `regular user`
	} else if loginCount > 10 {
		result = `Watch out`
	} else {
		result = `exactly 10`
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

}
