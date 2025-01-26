package main

import (
	"fmt"
)

func ReverseString(input string) string {
	// Konversi string ke slice rune untuk menangani karakter Unicode
	runes := []rune(input)
	n := len(runes)
	
	// Membalikkan slice rune
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	
	// Konversi kembali slice rune ke string
	return string(runes)
}

func main() {
	// Contoh penggunaan fungsi
	input := "Halo Dunia!"
	reversed := ReverseString(input)
	fmt.Println("String asli: ", input)
	fmt.Println("String terbalik: ", reversed)
}
