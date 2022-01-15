package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` // here - dosnt display this field in the json output
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to json video")
	EncodedJson()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func EncodedJson() {

	lcoCourses := []course{
		{"Reactjs Bootcam", 299, "lco.in", "abcd1234", []string{"web-dev", "js"}},
		{"Mern Bootcam", 199, "lco.in", "abc1234", []string{"full-stack", "js"}},
		{"Django Bootcap", 299, "lco.in", "abcd1234", nil},
	}

	// package this data as json data
	// MarshalIndent accptes 3 params (interface, prefix, intent)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	checkErr(err)
	fmt.Printf("%s\n", finalJson)
}
