package main

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func sha(data string, key string) string {
	//Create a new HMAC by defining the hash type and the key (as byte array)
	var h = hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))                 //Write data to it
	sha := hex.EncodeToString(h.Sum(nil)) // get result and encode as hexadecimal string
	return sha

}

func encrypt(publicKey rsa.PublicKey, msg string) string {
	cipher, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, []byte(msg), nil)
	data := string(cipher)
	return data
}

func decrypt(privateKey *rsa.PrivateKey, cipher []byte) []byte {
	plainText, _ := privateKey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA256})
	return plainText
}

func aliceDecrypts(cipherToSend string, secret string, privateKey *rsa.PrivateKey) {
	length := len(cipherToSend)
	mac := cipherToSend[length-64:]
	cipherReceived := cipherToSend[:length-64]
	hashMac := sha(cipherReceived, secret)
	fmt.Println("MAC", mac)
	fmt.Println("hashMac", hashMac)
	if mac == hashMac {
		fmt.Println(string(decrypt(privateKey, []byte(cipherReceived))))
	} else {
		fmt.Println("failed")

	}

}

func main() {
	secret := "mysecret"
	data := "tanvi"
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := privateKey.PublicKey
	cipher := encrypt(publicKey, data)
	sha := sha(cipher, secret)
	cipherToSend := string(cipher) + sha
	aliceDecrypts(cipherToSend, secret, privateKey)

}
