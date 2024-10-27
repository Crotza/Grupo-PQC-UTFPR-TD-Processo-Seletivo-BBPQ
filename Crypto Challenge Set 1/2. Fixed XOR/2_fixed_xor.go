package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func fixedXOR(buf1, buf2 []byte) ([]byte, error) {
	xorResult := make([]byte, len(buf1))

	for i := range buf1 {
		xorResult[i] = buf1[i] ^ buf2[i]
	}

	return xorResult, nil
}

func main() {
	hexString1 := "1c0111001f010100061a024b53535009181c"
	hexString2 := "686974207468652062756c6c277320657965"
	expectedResult := "746865206b696420646f6e277420706c6179"

	buf1, err := hex.DecodeString(hexString1)
	buf2, err := hex.DecodeString(hexString2)

	xorResult, err := fixedXOR(buf1, buf2)
	if err != nil {
		log.Fatal("Error performing XOR:", err)
	}

	xorHexString := hex.EncodeToString(xorResult)

	fmt.Println("Produced:", xorHexString)
	fmt.Println("Expected:", expectedResult)

	if xorHexString == expectedResult {
		fmt.Println("Correct answer!")
	} else {
		fmt.Println("Wrong answer!")
	}
}
