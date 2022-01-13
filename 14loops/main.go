package main

import (
	"fmt"
)

func main() {
	fmt.Println("Lopps in go")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d:=0; d < len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// if you dont want any value use _  eg say we replace index with _ it dosent return the value of index

	// for index, day := range days{
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}

		if rogueValue == 5 {
			break
		}

		fmt.Println("value is ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at Lco.in")

}
