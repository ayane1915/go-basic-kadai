package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

func saveDiary(entries []Entry) error {
	jsonBytes, marshalErr := json.Marshal(entries)
	if marshalErr != nil {
		// JSON変換エラー時の処理
		return marshalErr
	}
	writeErr := os.WriteFile("diary.json", jsonBytes, 0644)
	if writeErr != nil {
		// ファイル書き込みエラー時の処理
		return writeErr
	}
	return nil
}

func loadDiary() ([]Entry, error) {
	jsonBytes, readErr := os.ReadFile("diary.json")
	if os.IsNotExist(readErr) {
		// 初回はファイルがないため、空のスライスを返す
		return []Entry{}, nil
	}
	if readErr != nil {
		// ファイル読み込みエラー時の処理
		return nil, readErr
	}
	var entries []Entry
	unmarshalErr := json.Unmarshal(jsonBytes, &entries)
	if unmarshalErr != nil {
		// JSON変換エラー時の処理
		return nil, unmarshalErr
	}
	return entries, nil
}

func main() {
	// コマンドライン引数の数をチェック
	if len(os.Args) < 2 {
		fmt.Println("使い方：go run main.go [add|list] ...")
		return
	}

	// 1番目の引数（コマンド）を取得
	command := os.Args[1]

	// コマンドの内容に応じた処理
	switch command {
	case "add": // 日記エントリーの記録
		if len(os.Args) < 4 {
			fmt.Fprintln(os.Stderr, "使い方：go run main.go add <タイトル> <本文>")
			return
		}

		title := os.Args[2]
		content := os.Args[3]

		entries, err := loadDiary()
		if err != nil {
			fmt.Fprintln(os.Stderr, "日記の読み込みに失敗しました：", err)
			return
		}

		entries = append(entries, Entry{
			Timestamp: time.Now(),
			Title:     title,
			Content:   content,
		})
		if err := saveDiary(entries); err != nil {
			fmt.Fprintln(os.Stderr, "日記の保存に失敗しました：", err)
			return
		}

	case "list": // 日記の一覧表示
		entries, err := loadDiary()
		if err != nil {
			fmt.Fprintln(os.Stderr, "日記の読み込みに失敗しました：", err)
			return
		}
		for _, entry := range entries {
			fmt.Printf("日時: %s\nタイトル: %s\n本文: %s\n\n",
				entry.Timestamp.Format("2006-01-02 15:04:05"),
				entry.Title,
				entry.Content,
			)
		}

	default:
		fmt.Println("不明なコマンドです：", command)
	}
}
