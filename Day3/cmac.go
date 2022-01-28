package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/aead/cmac"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func genCmac(cipherText string, passphrase string) string {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	hashCmac, _ := cmac.New(block)
	hashCmac.Write([]byte(cipherText))
	hCmac := hex.EncodeToString(hashCmac.Sum(nil))
	return hCmac

}

func encrypt(data []byte, passphrase string) string {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)

	hCmac := genCmac(string(cipherText), passphrase)

	cipherToSend := string(cipherText) + hCmac

	return cipherToSend
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce := data[:nonceSize]
	cipherText := data[nonceSize:]
	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	return plainText

}

func checkAuthenticity(cipherToSend string, key string) {
	length := len(cipherToSend)
	c := cipherToSend[length-32:]
	cipherReceived := cipherToSend[:length-32]
	cmacObtained := genCmac(cipherReceived, key)
	fmt.Println("CMAC", c)
	fmt.Println("hashCmac", cmacObtained)

	if c == cmacObtained {
		fmt.Println(string(decrypt([]byte(cipherReceived), key)))
	} else {
		fmt.Println("failed")

	}

}

func main() {
	key := "hello"
	cipherText := encrypt([]byte("TanviShetty"), key)
	checkAuthenticity(cipherText, key)

}
