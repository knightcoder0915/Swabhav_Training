package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
)

func encrypts(msg string, publicKey rsa.PublicKey) []byte {
	cipher, _ := rsa.EncryptOAEP(sha512.New(), rand.Reader, &publicKey, []byte(msg), nil)
	return cipher

}

func decrypt(privateKey *rsa.PrivateKey, cipher []byte) []byte {
	plainText, _ := privateKey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA512})
	return plainText
}

func main() {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := privateKey.PublicKey
	msg := "tanvi"
	cipher := encrypts(msg, publicKey)
	fmt.Println("CipherText", string(cipher))
	plainText := decrypt(privateKey, cipher)
	fmt.Println("plainText", string(plainText))

}
