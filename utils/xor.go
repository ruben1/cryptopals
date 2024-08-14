package utils

func XorSingleCharacterKey(bytes []byte, key byte) []byte {
	result := make([]byte, len(bytes))
	for i:=0; i<len(bytes); i++ {
		result[i] = bytes[i] ^ key
	}
	return result
}

func XorRepeatingKey(bytes []byte, key []byte) []byte {
	result := make([]byte, len(bytes))
	for i:=0; i<len(bytes); i++ {
		result[i] = bytes[i] ^ key[i % len(key)]
	}
	return result
}