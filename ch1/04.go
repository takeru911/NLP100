package main

import (
    "fmt"
    "regexp"
)

func main(){
    str := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can"
	r := regexp.MustCompile(`[A-Za-z]+`)
	filteredStrs := r.FindAllStringSubmatch(str, -1)
    elementMap := make(map[int] string)

    for i, s := range filteredStrs {
        runeStr := []rune(s[0])

        if i == 0 || i == 4 || i == 5 || i == 6 || i == 7 || i == 8 || i == 14 || i == 15 || i == 18 {
            elementMap[i] = string(runeStr[0])
        }else{
            elementMap[i] = string(runeStr[0:2])
        }
    }

    for i := 0; i < len(elementMap); i++{
        fmt.Printf("element %d=%s\n", i, elementMap[i])
    }
 
    
}
