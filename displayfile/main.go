package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("quest8.txt")
	if err != nil {
		fmt.Printf("File name missing")
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arg")
	} else {
		fmt.Print(string(content))
	}
}
