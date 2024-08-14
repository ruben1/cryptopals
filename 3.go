package main

import (
	"encoding/hex"
	"fmt"
	"cryptopals/utils"
)

func decodeHex(hexStr string) ([]byte, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		fmt.Println("Error decoding hex:", err)
		return []byte{}, err
	}
	return bytes, nil
}

func main() {
	ciphertext, _ := decodeHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	
	candidates := utils.CrackXorCipher(ciphertext)
	// print top 5
	for i := 0; i < 5; i++ {
		fmt.Printf("Candidate key: %v, original input: %v, score: %v\n", candidates[i].Key, string(candidates[i].Plaintext), candidates[i].Score)
	}
}

// Lowest score:
// Candidate key: 88, original input: Cooking MC's like a pound of bacon, score: 0.6706941176470589
