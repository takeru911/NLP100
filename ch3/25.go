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
	r := regexp.MustCompile(`\{\{基礎情報.+\n([\s\S]*?)}}\n`)
	items := r.FindAllStringSubmatch(brText, -1)
	contents := items[0][1] + "\n|"
   	r = regexp.MustCompile(`(.+) = ([\s\S]+?)\n\|`)
	matchs := r.FindAllStringSubmatch(contents, -1)
	maps := map[string]string{}
	for _, match := range matchs {
		maps[match[1]] = match[2]
	}
    r = regexp.MustCompile(`'{2,5}`)
    matchLink := regexp.MustCompile(`\{\{`)
	for k, v := range maps {
        replace := r.ReplaceAllString(v, "")
        replace = matchLink.ReplaceAllString(replace, "")
		maps[k] = replace
        fmt.Println(maps[k])
	}

    
}
