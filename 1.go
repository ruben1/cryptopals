package main

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
)

func hexToBase64(hexStr string) ([]byte, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		fmt.Println("Error decoding hex:", err)
		return []byte{}, err
	}
	return bytes, nil
}

func main() {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	bytes, _ := hexToBase64(hexStr)
	fmt.Println(string(bytes))
	fmt.Println(base64.StdEncoding.EncodeToString(bytes))
}