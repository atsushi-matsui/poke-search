#!/bin/bash

# CSVファイルのパス
csv_file="./script/poke.csv"

# CSVファイルの内容を読み込んでAPIを呼び出す関数
call_api() {
    # 引数として受け取ったCSVファイルのパス
    local file=$1

    # CSVファイルの内容を処理するコードを記述
    # ここでは、CSVファイルを1行ずつ読み込んでAPIを呼び出す例とします
    while IFS=',' read -r id name type
    do
        # API呼び出しのコードを記述
        # ここでは、curlコマンドを使ってAPIを呼び出す例とします
        curl -X POST "http://localhost:8081/register/poke/$id?poke=$name"
        curl -X POST http://localhost:8081/register --data-urlencode id=$id --data-urlencode term=$type
    done < "$file"
}

# CSVファイルの読み込みとAPIの呼び出し
#call_api "$csv_file"
call_api $1
