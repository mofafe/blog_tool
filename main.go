package main

import (
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Html struct {
	Title    string
	Date     int
	Headline string
	Text     string
}

func Date() int {
	t := time.Now()
	format_t := t.Format("20060102")
	int_t, _ := strconv.Atoi(format_t)

	return int_t
}

func main() {
	d := Html{
		Title:    "test",
		Date:     Date(),
		Headline: "test",
		Text:     "test",
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
