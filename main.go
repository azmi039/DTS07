package main

import (
	"fmt"
)

func main() {
	kalimat := "selamat malam"
	huruf := make(map[string]int)

	for _, char := range kalimat {
		if string(char) == " " {
			fmt.Println()
		} else {
			fmt.Println(string(char))
			huruf[string(char)]++
		}
	}

	fmt.Println("map:")
	for hur, kem := range huruf {
		fmt.Printf("%s: %d", hur, kem)
	}
}