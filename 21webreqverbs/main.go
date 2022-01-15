package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to web verb")
	PerformGetRequest()
}
func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	checkErr(err)

	defer response.Body.Close()
	// fmt.Println("Status Code: ", response.StatusCode)
	// fmt.Println("Content Length: ", response.ContentLength)

	// method 1
	// content, err := ioutil.ReadAll(response.Body)
	// fmt.Println("content: ", string(content))

	// method 2
	var responseString strings.Builder

	content, err := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
