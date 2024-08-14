package utils

import (
	"sort"
	"math"
)

var letterFrequency = map[byte]float64{
	'a': 0.0817, 'b': 0.0149, 'c': 0.0278, 'd': 0.0425,
	'e': 0.1270, 'f': 0.0223, 'g': 0.0202, 'h': 0.0609,
	'i': 0.0697, 'j': 0.0015, 'k': 0.0077, 'l': 0.0403,
	'm': 0.0241, 'n': 0.0675, 'o': 0.0751, 'p': 0.0193,
	'q': 0.0010, 'r': 0.0599, 's': 0.0633, 't': 0.0906,
	'u': 0.0276, 'v': 0.0098, 'w': 0.0236, 'x': 0.0015,
	'y': 0.0197, 'z': 0.0007,
}
var letterByFrequency = []byte{101,116,97,111,105,110,115,114,104,108,100,99,117,109,102,112,103,119,121,98,118,107,120,106,113,122}

func xor(bytesA []byte, bytesB []byte) []byte {
	result := make([]byte, len(bytesA))
	for i:=0; i<len(bytesA); i++ {
		result[i] = bytesA[i] ^ bytesB[i]
	}
	return result
}

type kv struct {
	Key   byte
	Value int
}

func sortMap(charFreq map[byte]int) []byte {
	var ss []kv
	for k, v := range charFreq {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	sortedKeys := make([]byte, len(charFreq))
	for i, kv := range ss {
		sortedKeys[i] = kv.Key
	}

	return sortedKeys
}

func buildCharFreq(bytes []byte) map[byte]float64 {
	frequency := make(map[byte]float64)
	for _, b := range bytes {
		frequency[b]++
	}
	for k, v := range frequency {
		frequency[k] = v / float64(len(bytes))
	}
	return frequency
}

func scorePlaintextGuess(plaintext []byte) float64 {
	plaintextCharFreq := buildCharFreq(plaintext)
	// compare with common letter frequency
	score := 0.00
	for letter, expectedFrequency := range letterFrequency {
		actualFrequency, ok := plaintextCharFreq[letter]
		if !ok {
			actualFrequency = 0
		}
		score += math.Abs(expectedFrequency - actualFrequency)
	}
	return score
}

type Candidate struct {
    Key       byte
    Plaintext []byte
    Score     float64
}

func CrackXorCipher(ciphertext []byte) []Candidate {
	candidates := make([]Candidate, 256)
	
    for i := 0; i < 256; i++ {
		candidateKey := byte(i)
		xorOperand := make([]byte, len(ciphertext))
		for j := range xorOperand {
			xorOperand[j] = candidateKey
		}
		plaintext := xor(ciphertext, xorOperand)
		score := scorePlaintextGuess(plaintext)
		candidates[i] = Candidate{candidateKey, plaintext, score}
    }
	// sort by score
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Score < candidates[j].Score
	})
	return candidates
}
