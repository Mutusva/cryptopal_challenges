package main

/*

Fixed XOR
Write a function that takes two equal-length buffers and produces their XOR combination.

*/

import (
	"encoding/hex"
	"errors"
	"fmt"
)

const base = "686974207468652062756c6c277320657965"

func main()  {
	data := "1c0111001f010100061a024b53535009181c"

	result, err := FixedXOR(data, base)
	if err != nil {
		fmt.Println("could not perform the operation XOR")
	}
	fmt.Println(result)
}

func FixedXOR(a string, b string) (string, error) {

  aBytes, err := hex.DecodeString(a)

  if err != nil {
	  return "", errors.New("Cannot decode hex to string")
  }

  bBytes, err := hex.DecodeString(b)
  if err != nil {
	return "", errors.New("Cannot decode hex to string")
  }

  if len(aBytes) != len(bBytes) {
	return "", errors.New("The strings are not of the same length")
  }

  outBytes := make([]byte, len(aBytes))
  for i := 0; i < len(outBytes); i++ {
	outBytes[i] = aBytes[i]^bBytes[i]
  }

  return hex.EncodeToString(outBytes), nil
}