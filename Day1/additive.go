package main

import (
	"fmt"
	"strings"
)

func removeSpaces(plainText string) string {
	return strings.ReplaceAll(plainText, " ", "")
}

func encryption(str string, key int) string {
	plainText := removeSpaces(str)
	cipherText := ""
	for _, letter := range plainText {
		number := ((int(letter) - 97) + key) % 26
		char := string(number + 97)
		cipherText = cipherText + char
	}
	return cipherText
}

func decrypt(cipherText string, key int) string {
	plainText := ""
	for _, letter := range cipherText {
		number := ((int(letter) - 97) - key) % 26
		char := string(number + 97)
		plainText = plainText + char
	}
	return plainText
}

func main() {
	plainText := "Tanvi"
	key := 2

	cipherText := encryption(plainText, key)

	fmt.Println(cipherText)

	plainText1 := decrypt(cipherText, key)

	fmt.Println(plainText1)
}
