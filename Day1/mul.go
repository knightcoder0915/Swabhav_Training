package main

import (
	"fmt"
	"strings"
)

func removeSpaces(plainText string) string {
	return strings.ReplaceAll(plainText, " ", "")
}

func encrypt(plainText string, key int) string {
	pT := removeSpaces(plainText)
	cipherText := ""
	for _, letter := range pT {
		number := ((int(letter) - 97) * key) % 26
		char := string(number + 97)
		cipherText = cipherText + char
	}
	return cipherText
}

func inverse(key int) int {
	// a := GCD(key,26)
	// if a!=1{
	// 	return
	// }
	var i int
	for i = 0; i < 26; i++ {
		a := (key * i) % 26
		if a == 1 {
			return i
		}

	}
	return 0

}

func decrypt(cipherText string, key int) string {
	plainText := ""
	key = inverse(key)
	for _, letter := range cipherText {
		number := ((int(letter) - 97) * key) % 26
		char := string(number + 97)
		plainText = plainText + char
	}
	return plainText
}

func main() {
	plainText := "hello"
	key := 7

	cipherText := encrypt(plainText, key)

	fmt.Println(cipherText)

	plainText1 := decrypt(cipherText, key)

	fmt.Println(plainText1)
}
