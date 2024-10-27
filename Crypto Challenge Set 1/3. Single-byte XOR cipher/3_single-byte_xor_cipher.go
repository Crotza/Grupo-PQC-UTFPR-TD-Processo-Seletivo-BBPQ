package main

import (
	"encoding/hex"
	"fmt"
)

var frequencies = map[byte]float64{
	'a': 8.20, 'b': 1.50, 'c': 2.80, 'd': 4.30, 'e': 12.70, 'f': 2.20,
	'g': 2.00, 'h': 6.10, 'i': 7.00, 'j': 0.15, 'k': 0.77, 'l': 4.00,
	'm': 2.40, 'n': 6.70, 'o': 7.50, 'p': 1.90, 'q': 0.095, 'r': 6.00,
	's': 6.30, 't': 9.10, 'u': 2.80, 'v': 0.98, 'w': 2.40, 'x': 0.15,
	'y': 2.00, 'z': 0.074, ' ': 13.00,
}

func scoreText(text []byte) float64 {
	var score float64
	for _, b := range text {
		if b >= 'A' && b <= 'Z' {
			b += 'a' - 'A'
		}

		if freq, exists := frequencies[b]; exists {
			score += freq
		}
	}
	return score
}

func singleByteXOR(ciphertext []byte) ([]byte, byte, float64) {
	var bestText []byte
	var bestKey byte
	var bestScore float64

	for key := 0; key < 256; key++ {
		text := make([]byte, len(ciphertext))

		for i := range ciphertext {
			text[i] = ciphertext[i] ^ byte(key)
		}
		score := scoreText(text)

		if score > bestScore {
			bestScore = score
			bestText = text
			bestKey = byte(key)
		}
	}

	return bestText, bestKey, bestScore
}

func main() {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	ciphertext, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return
	}

	plaintext, key, score := singleByteXOR(ciphertext)
	fmt.Println("Key:", string(key))
	fmt.Println("Score:", score)
	fmt.Println("Produced:", string(plaintext))
}
