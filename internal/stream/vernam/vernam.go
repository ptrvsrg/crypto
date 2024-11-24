package vernam

import "crypto/rand"

func GenerateKey(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func Cipher(text, key []byte) []byte {
	result := make([]byte, len(text))
	for i := range text {
		result[i] = text[i] ^ key[i]
	}
	return result
}
