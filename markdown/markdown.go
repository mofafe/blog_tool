package markdown

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var buffer []string
var c bool
var s string

func MdToHtml() string {
	buffer = append(buffer, "\t\t\t<p>", "</p>\n")
	filename := "source.md"

	// ファイルオープン
	fp, err := os.Open(filename)
	if err != nil {
		// エラー処理
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var ht string

	for scanner.Scan() {
		// ここで一行ずつ処理
		s = scanner.Text()
		var m string
		if strings.HasPrefix(s, "#") {
			h := strings.Count(strings.Split(s, " ")[0], "#")
			if h <= 6 {
				if c {
					hashCount := strconv.Itoa(h)
					pt := strings.Join(buffer, " ")
					md := "\t\t\t<h" + hashCount + ">" + strings.TrimSpace(strings.TrimLeft(s, "#")) + "</h" + hashCount + ">\n"
					m = pt + md
				} else {
					hashCount := strconv.Itoa(h)
					m = "\t\t\t<h" + hashCount + ">" + strings.TrimSpace(strings.TrimLeft(s, "#")) + "</h" + hashCount + ">\n"
				}

			} else {
				//後で
			}
		} else {
			//後で
		}
		ht = ht + m
	}

	if err = scanner.Err(); err != nil {
		// エラー処理
	}
	return ht
}
