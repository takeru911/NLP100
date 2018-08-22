package main

import "fmt"

func main() {
	s := "パタトクカシーー"
	runes := []rune{}

	for i, r := range s {
		if i%2 == 1 {
			runes = append(runes, r)
		}
	}
	fmt.Println(string(runes))
}
