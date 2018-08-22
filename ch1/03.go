package main

import (
	"fmt"
	"regexp"
)

func WordCount(s string) int {
	r := []rune(s)

	return len(r)
}

func main() {
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	r := regexp.MustCompile(`[A-Za-z]+`)
	filteredStrs := r.FindAllStringSubmatch(str, -1)
	counter := make([]int, len(filteredStrs))

	for i, s := range filteredStrs {
		counter[i] = WordCount(s[0])
	}
	fmt.Println(counter)

}
