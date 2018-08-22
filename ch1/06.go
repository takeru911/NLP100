package main

import (
	"fmt"
	"strings"
)

func RemoveDuplicate(array []string) []string {
	encountered := map[string]bool{}
	removedArray := make([]string, 0, len(array))

	for _, elem := range array {
		if !encountered[elem] {
			encountered[elem] = true
			removedArray = append(removedArray, elem)
		}
	}

	return removedArray
}

func Union(x, y []string) []string {
	unionArray := append(x, y...)
	return RemoveDuplicate(unionArray)
}

func Intersection(x, y []string) []string {
	unionArray := append(x, y...)

	encountered := map[string]bool{}
	intersectArray := make([]string, 0, len(unionArray))
	for _, elem := range x {
		if !encountered[elem] {
			encountered[elem] = true
		}
	}
	for _, elem := range y {
		if encountered[elem] {
			intersectArray = append(intersectArray, elem)
		}
	}

	return RemoveDuplicate(intersectArray)
}

func main() {
	s1 := "paraparaparadise"
	s2 := "paragraph"

	s1_bigram, _ := ngram_char(s1, 2)
	s2_bigram, _ := ngram_char(s2, 2)
	fmt.Println(s1_bigram)
	fmt.Println(s2_bigram)
	unionArray := Union(s1_bigram, s2_bigram)
	intersectArray := Intersection(s1_bigram, s2_bigram)
	fmt.Println(unionArray)
	fmt.Println(intersectArray)
}

type NegativeNError int

func (e NegativeNError) Error() string {
	return fmt.Sprintf("n cannot negative value, your n is %v", e)
}

func ngram_char(s string, n int) ([]string, error) {
	if n < 0 {
		return nil, NegativeNError(n)
	}
	chars := []rune(strings.Replace(s, " ", "", -1))
	ngram := []string{}
	fmt.Println(string(chars))

	for i := 0; i <= len(chars)-n; i++ {
		tmp := []string{}
		for j := 0; j < n; j++ {
			tmp = append(tmp, string(chars[i+j]))
		}
		ngram = append(ngram, strings.Join(tmp, ""))
	}

	return ngram, nil
}
