package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"cryptopals/utils"
)

func main() {
	ciphertextbase64, _ := os.ReadFile("./input/7.txt")
	ciphertext, _ := base64.StdEncoding.DecodeString(string(ciphertextbase64))

	key := []byte("YELLOW SUBMARINE")
	plaintext, _ := utils.DecryptAesEcb(ciphertext, key)
	fmt.Println(string(plaintext))
}