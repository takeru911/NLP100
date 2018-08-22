package main

import (
    "fmt"
    "bufio"
    "os"
    "io"
    "strings"
)

func ReadFile(fileName string) []string {
    var fp *os.File
    lines := []string{}
    fp, err := os.Open(fileName)
    if err != nil{
        panic(err)
    }
    reader := bufio.NewReaderSize(fp, 4096)
    i := 0
    for {
        line, _, err := reader.ReadLine()
        if err == io.EOF {
            break
        }
        lines = append(lines, string(line))
        i += 1
    }

    return lines
}

func main(){
    lines := ReadFile("hightemp.txt")
    for _, line := range lines {
        fmt.Println(strings.Replace(line, "\t"," ", -1))
    }
}
