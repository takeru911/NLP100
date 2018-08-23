package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadFile(fileName string, numOfLine int) {
	var fp *os.File

	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(fp)
	i := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF || i > (numOfLine-1) {
			break
		}
		fmt.Println(string(line))
		i += 1
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println(len(os.Args))
		fmt.Println("14 $(filename) $(headLine)")
		os.Exit(1)
	}
	fileName := os.Args[1]
	headLine, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		fmt.Println("please headline is integer")
		os.Exit(1)
	}
	fmt.Printf("fileName: %v, headLine: %v\n", fileName, headLine)
	ReadFile(fileName, int(headLine))
}
