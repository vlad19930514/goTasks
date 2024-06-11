package main

import (
	"fmt"
	"os"
)

func main() {
	//go run main.go "Hello, its my page: http://localhost123.com See you"
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string argument.")
		return
	}
	fmt.Println(CheckLinks(os.Args[1]))
}

func CheckLinks(str string) string {
	url := []byte(str)
	values := []byte("http://")
	asterisk := []byte("*")[0]
	space := []byte(" ")[0]
	pointer := 0
	for i := range url {
		if pointer == len(values)-1 {
			if url[i] == space {
				pointer = 0
				continue
			}
			url[i] = asterisk
			continue
		}
		if url[i] == values[pointer] {
			pointer++
		} else {
			pointer = 0
		}
	}

	return string(url)
}
