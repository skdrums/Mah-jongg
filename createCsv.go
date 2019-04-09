package main

import (
    "log"
    "os"
)

func createCsv() {
    // ファイルを生成
    file, err := os.Create("./data/file.txt")
    if err != nil { // エラー処理
        log.Fatal(err)
    }
    // プログラムが終わったらファイルを閉じる
    defer file.Close()
}