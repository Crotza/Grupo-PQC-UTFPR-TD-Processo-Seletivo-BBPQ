package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedResult := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Error decoding hex:", err)
		return
	}

	base64String := base64.StdEncoding.EncodeToString(bytes)

	fmt.Println("Produced:", base64String)
	fmt.Println("Expected:", expectedResult)

	if base64String == expectedResult {
		fmt.Println("Correct answer!")
	} else {
		fmt.Println("Wrong answer!")
	}
}
