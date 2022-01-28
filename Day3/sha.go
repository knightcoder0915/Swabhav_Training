package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {

	secret := "mysecret"
	msg := "data"

	//Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(msg))                  //Write data to it
	sha := hex.EncodeToString(h.Sum(nil)) // get result and encode as hexadecimal string
	fmt.Println("Result", sha)
}
