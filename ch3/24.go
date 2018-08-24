package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Wiki struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func ReadJson(fileName string) []Wiki {
	var fp *os.File
	wikis := []Wiki{}

	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReaderSize(fp, 1000000)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		var wiki Wiki
		if err := json.Unmarshal(line, &wiki); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		wikis = append(wikis, wiki)
	}
	return wikis
}

func main() {
	wikiPages := ReadJson("jawiki-country.json")
	var brText string
	for _, wiki := range wikiPages {
		if wiki.Title == "イギリス" {
			brText = wiki.Text
			break
		}
	}
	r := regexp.MustCompile(`(ファイル:|File:)(.+?)\|(.+)`)
	items := r.FindAllStringSubmatch(brText, -1)
	for _, item := range items {
		fmt.Println(item[2])
	}
}
