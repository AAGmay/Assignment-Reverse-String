package main

import (
	"fmt"
)

func ReverseString(input string) string {
	runes := []rune(input)
	n := len(runes)
	
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	
	return string(runes)
}

func main() {
	input := "Halo Dunia!"
	reversed := ReverseString(input)
	fmt.Println("String asli: ", input)
	fmt.Println("String terbalik: ", reversed)
}
