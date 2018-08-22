package main

import "fmt"

func main() {
	s1 := "パトカー"
	s2 := "タクシー"
	r1 := []rune(s1)
	r2 := []rune(s2)
	concat := []rune{}

	for i := 0; i < len(r1); i++ {
		concat = append(concat, r1[i])
		concat = append(concat, r2[i])
	}

	fmt.Println(string(concat))

}
