package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
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

type prefecture struct {
	name string
	num  int
}

func Count(lines []string) map[string]int {
	countUp := map[string]int{}

	for _, line := range lines {
		countUp[line] += 1
	}

	return countUp
}

func main() {
	lines := ReadFile("hightemp.txt")
	col1 := make([]string, len(lines), len(lines))
	for i, line := range lines {
		col1[i] = line[0]
	}
	result := Count(col1)
	prefectures := make([]prefecture, len(result), len(result))
	i := 0
	for key, value := range result {
		prefectures[i] = prefecture{
			key,
			value,
		}
		i += 1
	}
	sort.Slice(prefectures, func(i, j int) bool {
		return prefectures[i].num > prefectures[j].num
	})
	for _, p := range prefectures {
		fmt.Println(p)
	}
}
