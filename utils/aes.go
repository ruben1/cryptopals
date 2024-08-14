package utils

import (
	"crypto/aes"
)

func DecryptAesEcb(ciphertext []byte, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	
	plaintext := make([]byte, len(ciphertext))
	for i, j:= 0, 16; i< len(ciphertext); i, j = i+16, j+16 {
		cipher.Decrypt(plaintext[i:j], ciphertext[i:j])
	}
	return plaintext, nil
}