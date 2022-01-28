package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key)) //storing key in slice of bytes
	//fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}

func decryption(data []byte, passphrase string) []byte {
	//key := []byte(createHash(passphrase))
	hashText := data[:32]
	//fmt.Println("Hash", string(hashText))
	encryptedText := data[32:]
	encryptHash := createHash(string(encryptedText))
	//fmt.Println("Encrypted Hash", encryptHash)
	if string(hashText) == encryptHash {
		key := []byte(createHash(passphrase))
		block, _ := aes.NewCipher(key)
		gcm, _ := cipher.NewGCM(block)
		nonceSize := gcm.NonceSize()
		nonce := encryptedText[:nonceSize]
		cipherText := encryptedText[nonceSize:]
		plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
		return plainText

	}
	return nil

}

func main() {
	cipherText := encrypt([]byte("TanviShetty"), "hello")
	hashText := createHash(string(cipherText))
	//fmt.Println(cipherText, "hash", hashText)
	hash := ""
	hash += hashText
	//fmt.Println("hash", hash)
	newCipher := hash + string(cipherText)
	fmt.Println(newCipher)
	plainText := decryption([]byte(newCipher), "hello")
	fmt.Println(string(plainText))

}
