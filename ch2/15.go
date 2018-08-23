package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadFile(fileName string, numOfTail int) {
	var fp *os.File
	tailLine := make([] string, numOfTail, numOfTail)
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(fp)
	i := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		tailLine[i % numOfTail] = string(line)
		i += 1
	}
	for _, line := range tailLine{
        fmt.Println(line)
    }
}

func main(){
	if len(os.Args) != 3 {
		fmt.Println(len(os.Args))
		fmt.Println("14 $(filename) $(tailLine)")
		os.Exit(1)
	}
	fileName := os.Args[1]
	tailLine, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		fmt.Println("please headline is integer")
		os.Exit(1)
	}
    ReadFile(fileName, int(tailLine))

}
