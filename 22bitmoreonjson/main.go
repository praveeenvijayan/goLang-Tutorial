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
	// EncodedJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Mern Bootcam",
		"Price": 199,
		"website": "lco.in",
		"tags": ["full-stack", "js"]
	}
	`)
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON Was Valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("json is not valid")
	}

	// some cases where you want to add data to a key value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("key is %v and value id %v and type is %T\n", key, value, value)
	}
}
