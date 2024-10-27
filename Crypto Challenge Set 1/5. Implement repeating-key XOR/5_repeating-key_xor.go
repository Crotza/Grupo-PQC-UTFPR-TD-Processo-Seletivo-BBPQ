package main

import (
	"encoding/hex"
	"fmt"
)

func repeatingKeyXOR(text, key string) string {
	textBytes := []byte(text)
	keyBytes := []byte(key)
	ciphertext := make([]byte, len(textBytes))

	for i := range textBytes {
		ciphertext[i] = textBytes[i] ^ keyBytes[i%len(keyBytes)]
	}

	return hex.EncodeToString(ciphertext)
}

func main() {
	text := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"

	encrypted := repeatingKeyXOR(text, key)

	expectedResult := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	fmt.Println("Produced:", encrypted)
	fmt.Println("Expected:", expectedResult)

	if encrypted == expectedResult {
		fmt.Println("Correct answer!")
	} else {
		fmt.Println("Incorrect answer!")
	}
}
