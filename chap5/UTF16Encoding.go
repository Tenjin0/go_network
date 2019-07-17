package main

import (
	"fmt"
	"unicode/utf16"
)

func UTF16Encoding() {

	str := "百度一下, 你就知道"

	runes := utf16.Encode([]rune(str))
	ints := utf16.Decode(runes)
	str = string(ints)

	fmt.Println("String length", len(runes))
	fmt.Println("Byte length", len(ints))
}
