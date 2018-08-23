package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"math"
	"strconv"
)

func ReadFile(fileName string) []string {
	var fp *os.File
	lines := []string{}
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(fp)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}
    return lines
}

func SplitStringLine(lines []string, splitSize int, isEven bool) [][]string {
    splitLines := make([][]string, splitSize, splitSize)
    numOfLine := len(lines)
    maxLine := int(math.Ceil(float64(numOfLine) / float64(splitSize)))
    fmt.Println(float64(numOfLine) / float64(splitSize))
    var split int
    for i, line := range lines {
        if isEven {
            split = i % splitSize
        }else {
            if ((split + 1) * maxLine) <= i {
                split += 1
            }
        }
        
        if splitLines == nil {
            splitLines[split] = []string{}
        }
        splitLines[split] = append(splitLines[split], line)
    }

    return splitLines
}

func main(){
    if len(os.Args) != 3{
        fmt.Println("16 $(fileName) $(splitSize)")
        os.Exit(1)
    }
    fileName := os.Args[1]
    splitSize,_ := strconv.ParseInt(os.Args[2], 10, 64)
    readLines := ReadFile(fileName)
    splitLines := SplitStringLine(readLines, int(splitSize), false)
    for i, splitLine := range splitLines {
        fmt.Printf("split: %v\n", i)
        for _, line := range splitLine{
            fmt.Println(line)
        }
    }
    
}
