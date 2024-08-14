package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
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
	input, err := ioutil.ReadFile("./input/4.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	globalCandidates := make([]utils.Candidate, 0)

	for _, line := range strings.Split(string(input), "\n") {
		ciphertext, _ := decodeHex(line)
		localCandidates := utils.CrackXorCipher(ciphertext)
		globalCandidates = append(globalCandidates, localCandidates[0])
	}

	// sort by score
	sort.Slice(globalCandidates, func(i, j int) bool {
		return globalCandidates[i].Score < globalCandidates[j].Score
	})
	// print first 5
	for i := 0; i < 5; i++ {
		fmt.Printf("Plaintext: %v, Candidate key: %v, Score: %v\n", string(globalCandidates[i].Plaintext), globalCandidates[i].Key, globalCandidates[i].Score)
	}
}

// Plaintext: Now that the party is jumping, Candidate key: 53, score: 0.5918666666666668
