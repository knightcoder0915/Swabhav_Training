package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key)) //storing key in slice of bytes
	//fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	hashText := createHash("TanviShetty")
	fmt.Println(hashText)

}
