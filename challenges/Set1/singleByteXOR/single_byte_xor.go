package main

import (
	"fmt"
	"encoding/hex"
)

func main()  {
	txt := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	SingleByteXOR(txt)
	// Answer is Cooking MC's like a pound of bacon, and key is x(88)
}

func SingleByteXOR(text string) {
	b, _ := hex.DecodeString(text)

	for i := 0; i <= 256; i++ {
		bRes := make([]byte, len(b))
		for j,v := range b {
            bRes[j] = v ^ byte(i)
		}
		fmt.Printf("key=%v, and text=%s\n",i,string(bRes))
	}
	
}