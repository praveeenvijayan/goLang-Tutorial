package main

import "fmt"

// Public
const LoginToken string = "Hellooooooooos!"

func main() {
	var username string = "Praveen"
	fmt.Println(username)
	fmt.Printf("Variable is of type:  %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type:  %T \n", isLoggedIn)

	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type:  %T \n", smallVal)

	var smallFloat float32 = 255.455555534434454434
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type:  %T \n", smallFloat)

	var smallFloats float64 = 255.455555534434454434
	fmt.Println(smallFloats)
	fmt.Printf("Variable is of type:  %T \n", smallFloats)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type:  %T \n", anotherVariable)

	// implicit
	var website = "lco.in"
	fmt.Println(website)

	// no var style (valarus operator can be used inside methods)
	numberOfUsers := 30000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type:  %T \n", LoginToken)

}

