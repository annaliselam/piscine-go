package main

import (
	"fmt"
	"os"
	//	"io/ioutil"
)

func main() {
	multipleArgs := len(os.Args)
	if multipleArgs > 2 {
		fmt.Println("Too many arguments")
		return
	} else if multipleArgs < 2 {
		fmt.Println("File name missing")
		return
	}

	file, err := os.Open("quest8.txt")

	arr := make([]byte, 14)
	file.Read(arr)

	fmt.Println(string(arr))

	file.Close()
}
