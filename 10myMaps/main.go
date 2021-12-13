package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	// map[KeyNAme]DataType
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["py"] = "Python"
	languages["go"] = "GO Lang"

	fmt.Println("List of all laguages: ", languages)
	fmt.Println("Js shorts for: ", languages["JS"])

	delete(languages, "py")
	fmt.Println("List of all laguages: ", languages)

	// loops

	for key, value := range languages {
		fmt.Printf("For key %v, and Value is %v\n", key, value)
	}
}
