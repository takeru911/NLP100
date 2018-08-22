package main

import (
	"fmt"
	"strings"
)

type NegativeNError int

func (e NegativeNError) Error() string {
	return fmt.Sprintf("n cannot negative value, your n is %v", e)
}

func ngram_word(s string, n int) ([]string, error) {
	if n < 0 {
		return nil, NegativeNError(n)
	}

	words := strings.Fields(s)
	ngram := []string{}

	for i := 0; i <= len(words)-n; i++ {
		tmp := []string{}
		for j := 0; j < n; j++ {
			tmp = append(tmp, words[i+j])
		}
		ngram = append(ngram, strings.Join(tmp, " "))
	}
	return ngram, nil
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
		ngram = append(ngram, strings.Join(tmp, " "))
	}

	return ngram, nil
}

func main() {
	str := "I am an NLPer"
	word_result, ok := ngram_word(str, 2)
	if ok == nil {
		fmt.Println(word_result)
	}
	char_result, ok := ngram_char(str, 5)
	if ok == nil {
		fmt.Println(char_result)
	}

}
