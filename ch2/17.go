package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadFile(fileName string) [][]string {
	var fp *os.File
	fp, err := os.Open(fileName)
	defer fp.Close()
	r := csv.NewReader(fp)
	r.Comma = '\t'
	lines, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	return lines
}

func RemoveDuplicate(lines []string) []string {
	removedLine := make([]string, 0, len(lines))
	existElem := map[string]bool{}

	for _, line := range lines {
		if !existElem[line] {
			removedLine = append(removedLine, line)
			existElem[line] = true
		}
	}

	return removedLine
}

func main() {
	lines := ReadFile("hightemp.txt")
	col1 := make([]string, len(lines), len(lines))
	for i, line := range lines {
		col1[i] = line[0]
	}
	result := RemoveDuplicate(col1)
	for _, row := range result {
		fmt.Println(row)
	}
}
