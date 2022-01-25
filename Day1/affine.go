package main

import (
	"fmt"
	"strings"
)

func removeSpaces(plainText string) string {
	return strings.ReplaceAll(plainText, " ", "")
}

func encrypt(plainText string, a, b int) string {
	pT := removeSpaces(plainText)
	cipherText := ""
	for _, letter := range pT {
		number := (((int(letter) - 97) * a) + b) % 26
		char := string(number + 97)
		cipherText = cipherText + char
	}
	return cipherText
}

func inverse(key int) int {

	var i int
	for i = 0; i < 26; i++ {
		a := (key * i) % 26
		if a == 1 {
			return i
		}

	}
	return 0

}

func decrypt(cipherText string, a, b int) string {
	plainText := ""
	a = inverse(a)
	for _, letter := range cipherText {
		number := (((int(letter) - 97) - b) * a) % 26
		char := string(number + 97)
		plainText = plainText + char
	}
	return plainText
}

func main() {
	plainText := "hello"
	a, b := 11, 7

	cipherText := encrypt(plainText, a, b)

	fmt.Println(cipherText)

	plainText1 := decrypt(cipherText, a, b)

	fmt.Println(plainText1)
}
