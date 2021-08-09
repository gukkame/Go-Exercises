package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("quest8.txt")
	if err != nil || len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		fmt.Print(string(content))
	}
}
