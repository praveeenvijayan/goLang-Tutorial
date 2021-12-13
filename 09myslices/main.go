package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sclices")
	// var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("type of fruitlist is %T\n", fruitList)
	// fmt.Println("type of fruitlist is \n", fruitList)

	// fruitList = append(fruitList, "Orange", "Mango")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	// highScores := make([]int, 4)

	// highScores[0] = 234
	// highScores[1] = 945
	// highScores[2] = 465
	// highScores[3] = 867
	// // highScores[4] = 777

	// highScores = append(highScores, 555, 666, 321)

	// fmt.Println(highScores)

	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove a value  from slices based on iNdex

	var courses = []string{"React", "Python", "go", "Flutter"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
