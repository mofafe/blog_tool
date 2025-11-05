package markdown

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MdToHtml() string {
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
		s := scanner.Text()
		var hashLast bool = true
		var h string
		var i int

		if strings.Contains(s, "#") {
			for hashLast {
				hash := string(s[i])
				if hash == "#" {
					h = h + hash
				} else {
					hashLast = false
				}
				i++
			}
		}
		hashCount := strings.Count(h, "#")
		var m string
		switch hashCount {
		case 0:
			//後で書く
		case 1:
			m = "<h1>" + strings.TrimSpace(strings.TrimLeft(s, "#")) + "</h1>\n"
		case 2:
			m = "<h2>" + strings.TrimSpace(strings.TrimLeft(s, "#")) + "</h2>\n"
		case 3:
			m = "<h3>" + strings.TrimSpace(strings.TrimLeft(s, "#")) + "</h3>\n"
		case 4:
			m = "<h4>" + strings.TrimSpace(strings.TrimLeft(s, "#")) + "</h4>\n"
		case 5:
			m = "<h5>" + strings.TrimSpace(strings.TrimLeft(s, "#")) + "</h5>\n"
		case 6:
			m = "<h6>" + strings.TrimSpace(strings.TrimLeft(s, "#")) + "</h6>\n"
		}
		ht = ht + m
	}

	if err = scanner.Err(); err != nil {
		// エラー処理
	}
	fmt.Println(ht)
	return ht
}
