package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadTsvFile(fileName string) [][]string {
	var fp *os.File
	lines := [][]string{}
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReaderSize(fp, 4096)
	i := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		rawLine := string(line)
		columns := strings.Split(rawLine, "\t")
		lines = append(lines, columns)

		i += 1
	}

	return lines
}

func main() {
	lines := ReadTsvFile("hightemp.txt")
	col1 := []string{}
	col2 := []string{}

	for _, line := range lines {
		col1 = append(col1, line[0])
		col2 = append(col2, line[1])
	}
    ioutil.WriteFile("col1.txt", []byte(strings.Join(col1, "\n")), os.ModePerm)
    ioutil.WriteFile("col2.txt", []byte(strings.Join(col2, "\n")), os.ModePerm)    
}
