package main

/*
Convert hex to base64

The string(hex represenatation of): 45df

Should produce:
Rd8=
*/

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main()  {
	data := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	
	fmt.Println(ConvertHexToBase64(data))
}

func ConvertHexToBase64(data string) string {

	hex_s , err := hex.DecodeString(data)
	if err != nil {
		fmt.Println("could not convert to hex!")
	}
	return base64.StdEncoding.EncodeToString([]byte(hex_s))
}