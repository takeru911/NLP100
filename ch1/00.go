package main

import "fmt"

func main() {
	s := "stressed"
	reversed := make([]rune, len(s))

	for i, r := range s {
		reversed[len(reversed)-(i+1)] = r
	}
	fmt.Println(s)
	fmt.Println(string(reversed))
}
