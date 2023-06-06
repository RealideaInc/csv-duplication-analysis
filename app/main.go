package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	file, err := os.Open("file/ファイル名を入れてください")
	if err != nil {
		fmt.Println("ファイルを開くのに失敗しました:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))

	dataMap := make(map[string]int)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("CSVの読み込みに失敗しました:", err)
				return
			}
		}

		column1 := record[0]
		column2 := record[1]
		column8 := record[7]

		// 1番目、2番目、3番目のカラムの値をキーとしてマップに記録
		key := fmt.Sprintf("%s-%s-%s", column1, column2, column8)
		dataMap[key]++
	}

	// 完全に一致する組み合わせを抽出
	var result [][]string
	for key, count := range dataMap {
		if count > 1 {
			// 重複している組み合わせのみ結果に追加
			combination := strings.Split(key, "-")
			result = append(result, combination)
		}
	}

	// 結果を表示
	for _, combination := range result {
		fmt.Println(combination)
	}
	outputFile, err := os.Create("file/output.csv")
if err != nil {
    fmt.Println("ファイルの作成に失敗しました:", err)
    return
}
defer outputFile.Close()

// CSVライターを作成
writer := csv.NewWriter(outputFile)

// 結果をCSVファイルに書き込む
for _, combination := range result {
    if err := writer.Write(combination); err != nil {
        fmt.Println("ファイルの書き込みに失敗しました:", err)
        return
    }
}

// バッファの内容をフラッシュして、CSVファイルに書き込みを完了
writer.Flush()

if err := writer.Error(); err != nil {
    fmt.Println("ファイルの書き込みに失敗しました:", err)
    return
}

fmt.Println("結果をファイルに書き込みました。")
}
