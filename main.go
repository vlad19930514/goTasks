package main

import (
	"fmt"
)

func main() {
	url := "http://localhost123.com "

	replaceUtl := replaceUtlT{
		replaceValue: "*",
		startsWith:   "http://",
	}

	fmt.Println(replaceUtl.Replace(url))

}

type replaceUtlT struct {
	replaceValue string
	startsWith   string
}

func (data replaceUtlT) compareBites(first []byte, second []byte) bool {
	for i := range second {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func (data replaceUtlT) Replace(urlValue string) string {

	replaceValue := []byte(data.replaceValue)[0]
	url := []byte(urlValue)
	start := []byte(data.startsWith)

	shouldBeReplaced := data.compareBites(url[:len(start)], start)

	if shouldBeReplaced {
		result := make([]byte, 0, len(url))
		result = append(result, start...)
		replaceFrom := len(start) - 1

		for i := replaceFrom; i < len(url); i++ {
			result = append(result, replaceValue)
		}
		return string(result)
	}

	return urlValue
}
