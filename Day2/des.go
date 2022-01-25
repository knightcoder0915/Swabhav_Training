package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {

	triplekey := "876543218765432187654321" //24bytes for three keys

	plaintext := []byte("Hello Wo") // Hello Wo = 8 bytes.

	block, _ := des.NewTripleDESCipher([]byte(triplekey))
	/*if err != nil {
		fmt.Printf("%s \n", err.Error())
		os.Exit(1)
	}*/

	fmt.Printf("%d bytes NewTripleDESCipher key with block size of %d bytes\n", len(triplekey), block.BlockSize)

	ciphertext := []byte("abcdef1234567890")
	iv := ciphertext[:des.BlockSize] // const BlockSize = 8

	// encrypt

	mode := cipher.NewCBCEncrypter(block, iv)
	//encrypted := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext[des.BlockSize:], plaintext)
	//mode.CryptBlocks(encrypted, plaintext)
	fmt.Printf("%s encrypt to %x \n", plaintext, ciphertext[des.BlockSize:])

	//decrypt
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(plaintext))
	decrypter.CryptBlocks(decrypted, ciphertext[des.BlockSize:])
	fmt.Printf("%x decrypt to %s\n", ciphertext[des.BlockSize:], decrypted)

}
