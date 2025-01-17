package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LocalisableStringsAnalyzer(path string, infos *[]AnalyzedInfrmation) {

	// xxx.strings のみを解析
	if !strings.HasSuffix(path, ".strings") {
		return
	}

	// ファイルをOpenする
	file, err := os.Open(path)
	// 読み取り時の例外処理
	if err != nil {
		fmt.Println("error")
	}
	// 関数が終了した際に確実に閉じるようにする
	defer file.Close()

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}

		text := sc.Text()
		text = strings.TrimSpace(text)
		// 後ろにコメントが付いていれば削除する。
		index := strings.Index(text, "//")
		if index >= 0 {
			text = text[:index]
			text = strings.TrimSpace(text)
		}

		if strings.HasSuffix(text, ";") {
			index := strings.Index(text, "=")
			if index >= 0 {
				text = text[:index]
				text = strings.TrimSpace(text)
				text = strings.Trim(text, "\"")
			}

			var result bool = false
			for _, element := range *infos {
				if element.Description == text {
					result = true
					break
				}
			}
			if result == false {
				//*texts = append(*texts, text)
				*infos = append(*infos, AnalyzedInfrmation{path, 0, text})
			}
		}
	}
}
