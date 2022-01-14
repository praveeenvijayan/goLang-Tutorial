package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("working with files")

	content := "This needs to go in a file - Lco.in"

	file, err := os.Create("./mylcogofile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	// databyte give err in bytes so always convert it to string 
	fmt.Println("text data inside the file is \n", string(databyte))
}
func checkNilErr(err error)  {
	if err != nil {
		panic(err)
	}
}