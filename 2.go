package main

import (
	"encoding/hex"
	"fmt"
)

func decodeHex(hexStr string) ([]byte, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		fmt.Println("Error decoding hex:", err)
		return []byte{}, err
	}
	return bytes, nil
}

func xor(bytesA []byte, bytesB []byte) []byte {
	result := make([]byte, len(bytesA))
	for i:=0; i<len(bytesA); i++ {
		result[i] = bytesA[i] ^ bytesB[i]
	}
	return result
}

func main() {
	bytesA, _ := decodeHex("1c0111001f010100061a024b53535009181c")
	bytesB, _ := decodeHex("686974207468652062756c6c277320657965")
	bytes := xor(bytesA, bytesB)
	fmt.Println(string(bytes))
	fmt.Println(hex.EncodeToString(bytes))
}