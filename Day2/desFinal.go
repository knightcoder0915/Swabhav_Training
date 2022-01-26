package main

import (
	"bufio"
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"os"
)

func DesEncryption(key, iv, plainText []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

func DesDecryption(key, iv, cipherText []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
func GenerateByteKey(s string) []byte {
	key := []byte(s)
	if len(key) > 8 {
		fmt.Println("len of byte greater than 8  0-7")
		key = key[:8]
		return key
	} else if len(key) < 8 {
		return PKCS5Padding(key, 8)
	}
	return key
}
func GenerateIVByte(s string) []byte {
	iv := []byte(s)
	if len(iv) > 8 {
		iv = iv[:8]
		return iv
	} else if len(iv) < 8 {
		return PKCS5Padding(iv, 8)
	}
	return iv
}
func main() {
	var originalText string
	fmt.Println("Enter original text")
	reader := bufio.NewReader(os.Stdin)
	originalText, _ = reader.ReadString('\n')
	fmt.Println(originalText)
	mytext := []byte(originalText)
	var key string
	fmt.Println("Enter your key")
	fmt.Scanln(&key)
	keyByte := GenerateByteKey(key)
	fmt.Println("Key Byte ", keyByte)
	//fmt.Println("Enter initialization vector ")
	var iv string = "129204679"
	//fmt.Scanln(&iv)
	ivByte := GenerateIVByte(iv)
	//ivByte := []byte("43218765")
	cryptoText, _ := DesEncryption(keyByte, ivByte, mytext)
	fmt.Println(string(cryptoText))
	decryptedText, _ := DesDecryption(keyByte, ivByte, cryptoText)
	fmt.Println(string(decryptedText))

}
