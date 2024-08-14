package main

import (
	"encoding/base64"
	"math/bits"
	"fmt"
	"os"
	"sort"
	"cryptopals/utils"
	"gonum.org/v1/gonum/stat/combin"
)

type kv struct {
	Key   int
	Value float64
}

func editDistance(inputA []byte, inputB []byte) int {
	count := 0
	for index, byteA := range inputA {
		xor := byteA ^ inputB[index]
		count += bits.OnesCount8(xor)

	}
	return count
}

func getKeySizeScore(ciphertext []byte, keysize int) float64 {
	chunks := [][]byte{
		ciphertext[0:keysize],
		ciphertext[keysize:keysize*2],
		ciphertext[keysize*2:keysize*3],
		ciphertext[keysize*3:keysize*4],
	}
	list := combin.Combinations(4, 2)
	distance := 0.0
	for _, v := range list {
		distance += float64(editDistance(chunks[v[0]], chunks[v[1]]))
	}
	distance /= float64(len(list))
	return distance / float64(keysize)
}

func main() {
	ciphertextbase64, _ := os.ReadFile("./input/6.txt")
	ciphertext, _ := base64.StdEncoding.DecodeString(string(ciphertextbase64))

	keySizeCandidates := make([]kv, 0)

	for i := range 40 {
		keysize := i+1
		keySizeScore := getKeySizeScore(ciphertext, keysize)
		keySizeCandidates = append(keySizeCandidates, kv{keysize, keySizeScore})
	}

	sort.Slice(keySizeCandidates, func(i, j int) bool {
		return keySizeCandidates[i].Value < keySizeCandidates[j].Value
	})

	keySizes := keySizeCandidates[:3]
	fmt.Println(keySizes)
	for _,item := range keySizes {
		keysize := item.Key
		blocks := make([][]byte, keysize)
		for i:=0; i<len(ciphertext); i++ {
			blocks[i % keysize] = append(blocks[i % keysize], ciphertext[i])
		}
		
		key := make([]byte, keysize)

		plaintextscore := 0.0
		for idx, block := range blocks {	
			plaintext := utils.CrackXorCipher(block)
			plaintextscore += plaintext[0].Score
			key[idx] = plaintext[0].Key
		}
		plaintextscore /= float64(keysize)
		fmt.Println("Keysize:", keysize)
		fmt.Println("Plaintext score:", plaintextscore)
		fmt.Println("Key:", string(key))
		fmt.Println("Plaintext:", string(utils.XorRepeatingKey(ciphertext, key)))
	}
}