package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

func bytesToChunks(input []byte, blockSize int) [][]byte {
	chunks := make([][]byte, 0)
	for i := 0; i < len(input); i += blockSize {
		end := i + blockSize
		if end > len(input) {
			end = len(input)
		}
		chunks = append(chunks, input[i:end])
	}
	return chunks
}

func main() {
	ciphertextHex, err := ioutil.ReadFile("./input/8.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	for i, line := range strings.Split(string(ciphertextHex), "\n") {
		ciphertext, err := hex.DecodeString(line)
		if err != nil {
			fmt.Println("Error decoding hex:", err)
			return
		}
		chunks := bytesToChunks(ciphertext, 16)

		uniqueChunks := make(map[string]bool)
		for _, chunk := range chunks {
			uniqueChunks[string(chunk)] = true
		}
		if (len(chunks) != len(uniqueChunks)) {
			fmt.Println("Found ECB at line", i)
			fmt.Printf("Chunk length %v, Unique chunk length: %v\n", len(chunks), len(uniqueChunks))
		}
	}
}