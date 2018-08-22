package main

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"strings"
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

func main() {
	col1 := ReadFile("col1.txt")
	col2 := ReadFile("col2.txt")
	result := []string{}
	for i := 0; i < len(col1); i++ {
		record := strings.Join([]string{col1[i][0], col2[i][0]}, "\t")
		result = append(result, record)
	}
	ioutil.WriteFile("merge.txt", []byte(strings.Join(result, "\n")), os.ModePerm)
}
