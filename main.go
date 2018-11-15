package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(WordReversal("Australia is a country and continent surrounded by the Indian and Pacific oceans"))
	fmt.Println(LetterReversal("Australia is a country and continent surrounded by the Indian and Pacific oceans"))
}

func WordReversal(sentence string) string {
	outputString := ""
	testArray := strings.Fields(sentence)
	for i := len(testArray) - 1; i >= 0; i-- {
		s := string(testArray[i])
		outputString += " " + s
	}
	return strings.TrimSpace(outputString)
}

func LetterReversal(sentence string) string {
	outputString := ""
	testArray := strings.Split(sentence, "")
	for i := len(testArray) - 1; i >= 0; i-- {
		s := string(testArray[i])
		outputString += s
	}
	return strings.TrimSpace(outputString)
}
