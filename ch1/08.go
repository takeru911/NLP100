package main

import (
    "fmt"
    "strings"
)

func ciper(s string) string{
    decoded := []string{}
    r := []rune(s)

    for _, c := range r {
        if c >= 97 && c <= 122 {
            decoded = append(decoded, string(219 - c))
        }else{
            decoded = append(decoded, string(c))
        }
    }
    return strings.Join(decoded, "")
}

func main(){
    s := "takeru is maehara 3"
    fmt.Println(s)
    fmt.Println(ciper(s))
}
