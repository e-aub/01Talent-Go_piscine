package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) != 2 {
		fmt.Println("Too many arguments")
	} else {
		filename := string(os.Args[1])
		filecontent, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("There is no such file in this directory")
		} else {
			fmt.Println(string(filecontent))
		}
	}
}
