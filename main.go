package main

import (
	"fmt"
	"log"
	"mofafe/blog_tool/markdown"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Html struct {
	Title       string
	Date        int
	Description string
	Text        string
}

func Date() int {
	t := time.Now()
	format_t := t.Format("20060102")
	int_t, _ := strconv.Atoi(format_t)

	return int_t
}

func main() {
	var title string
	var description string
	fmt.Scan(&title)
	fmt.Scan(&description)

	d := Html{
		Title:       title,
		Date:        Date(),
		Description: description,
		Text:        markdown.MdToHtml(),
	}

	DateString := strconv.Itoa(Date())

	file, err := os.Create(DateString + ".html")
	if err != nil {
		log.Fatal(err) // エラー処理
	}
	defer file.Close() // 処理終了後にファイルを閉じる

	tpl := template.Must(template.ParseFiles("template.html"))
	tpl.Execute(file, d)
}
