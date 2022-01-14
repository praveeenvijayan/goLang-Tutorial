package main

import "fmt"

func main() {
	fmt.Println("Functins in Go")
	greeter()
	result := adder(3, 4)
	fmt.Println("result is : ", result)

	proRes, myMsg := proAdder(2, 3, 4, 5)
	fmt.Println("result is : ", proRes, myMsg)

}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}


// retuning mutliple datatypes
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val

	}
	return total, "hey there its my message"
}

func greeter() {
	fmt.Println("Namasthea From Go")
}
