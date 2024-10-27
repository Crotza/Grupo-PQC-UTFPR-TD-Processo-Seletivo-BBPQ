package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
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
	file, err := os.Open("single-character_xor.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var bestText []byte
	var bestKey byte
	var bestScore float64
	var bestLine string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		ciphertext, err := hex.DecodeString(line)
		if err != nil {
			fmt.Println("Error decoding hex string:", err)
			continue
		}

		text, key, score := singleByteXOR(ciphertext)

		if score > bestScore {
			bestScore = score
			bestText = text
			bestKey = key
			bestLine = line
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Encrypted line:", string(bestLine))
	fmt.Println("Key:", string(bestKey))
	fmt.Println("Score:", bestScore)
	fmt.Println("Decrypted text:", string(bestText))
}
