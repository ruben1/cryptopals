package main

import "fmt"

func pkcs7(input []byte, blocksize int) []byte {
	padLength := blocksize - (len(input) % blocksize)
	for i:=0; i<padLength; i++ {
		input = append(input, byte(padLength))
	}
	return input
}

func unpkcs7(input []byte) (error, []byte) {
	padLength := int(input[len(input)-1])

	if padLength == 0 || padLength > len(input) {
		return fmt.Errorf("Invalid padding"), nil
	}

	for i:=len(input)-1; i>len(input)-padLength; i-- {
		if input[i] != byte(padLength) {
			return fmt.Errorf("Invalid padding"), nil
		}
	}
	return nil, input[:len(input)-padLength]
}

func main() {
	input := []byte("YELLOW SUBMARINE")
	paddedInput := pkcs7(input, 20)
	_, unpaddedInput := unpkcs7(paddedInput)
	fmt.Println("Original input:", input)
	fmt.Println("Padded input:", paddedInput)
	fmt.Println("Unpadded input:", unpaddedInput)
}
