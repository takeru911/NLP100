package main

import (
    "fmt"
    "regexp"
    "math/rand"
)

func main(){
    s := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
    r := regexp.MustCompile(`[A-Za-z']+`)
    filteredStrs := r.FindAllStringSubmatch(s, -1)
    results := []string{}

    for _, word := range filteredStrs {
        runes := []rune(word[0])

        if len(runes) <= 4 {
            results = append(results, string(runes))
        }else {
            randomRune := make([]rune, 0, len(runes))
            randomRune = append(randomRune, runes[0])
            perm := rand.Perm(len(runes) - 2)
            for _, v := range perm {
                randomRune = append(randomRune, runes[v + 1])
            }
            randomRune = append(randomRune, runes[len(runes) - 1])
            results = append(results, string(randomRune))
        }
    }
    fmt.Println(results)
}
