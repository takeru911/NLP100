package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
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

func ConvertRow(lines [][]string) []Row {
	rows := make([]Row, len(lines), len(lines))
	dateLayout := "2006-01-02"
	for i, line := range lines {
		temp, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			fmt.Println("col3 is not float")
			os.Exit(1)
		}
		date, err := time.Parse(dateLayout, line[3])
		if err != nil {
			fmt.Println("col4 is not yyyy-mm-dd format")
			os.Exit(1)
		}
		rows[i] = Row{
			line[0],
			line[1],
			temp,
			date,
		}
	}
	return rows
}

type Row struct {
	prefecture string
	city       string
	temp       float64
	date       time.Time
}

func (r *Row) String() string {
	return fmt.Sprintf("prefecture: %s, city: %s, temp: %v, date: %v", r.prefecture, r.city, r.temp, r.date)
}

func main() {
	lines := ReadFile("hightemp.txt")
	rows := ConvertRow(lines)
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].date.Sub(rows[j].date) > 0
	})
	for _, row := range rows {
		fmt.Println(&row)
	}
}
