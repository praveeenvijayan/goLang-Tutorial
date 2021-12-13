package main

import "fmt"

func main() {
	fmt.Println("Welcome toarray in golang: ")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Orange"

	fmt.Println("FruitList is", fruitList)
	fmt.Println("Length of FruitList is", len(fruitList))


	var vegList = [3] string {"Carrot", "Beans", "Onion"}

	fmt.Println("Vegie List is", vegList)
	fmt.Println("Vegie List is", len(vegList))
}
