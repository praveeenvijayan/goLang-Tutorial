package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to web verb")
	// PerformGetRequest()
	PerformPostJsonRequest()
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

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json pay load

	requestBody := strings.NewReader(`
		{
			"coursename":"Lets go with go lang",
			"price": 0,
			"platform": "lco.in"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	checkErr(err)
	content, err := ioutil.ReadAll(response.Body)
	fmt.Println("content: ", string(content))
	checkErr(err)
	defer response.Body.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
