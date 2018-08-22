package main

import (
    "fmt"
    "bufio"
    "os"
    "io"
)

func main(){
    var fp *os.File
    fp, err := os.Open("hightemp.txt")
    if err != nil{
        panic(err)
    }
    reader := bufio.NewReaderSize(fp, 4096)
    i := 0
    for {
        _, _, err := reader.ReadLine()
        if err == io.EOF {
            break
        }
        i += 1
    }
    fmt.Println(i)
}
