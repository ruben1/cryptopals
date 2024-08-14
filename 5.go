package main

import (
	"encoding/hex"
	"fmt"
	"cryptopals/utils"
)

func main() {
	key := []byte("ICE")
	plaintext := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")

	ciphertext := utils.XorRepeatingKey(plaintext, key)
	fmt.Println(hex.EncodeToString(ciphertext))
}