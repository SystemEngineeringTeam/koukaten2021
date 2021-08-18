package apifuncs

import (
	"encoding/json"
	"fmt"
	"net/http"

	"set1.ie.aitech.ac.jp/koukaten2021/db"
)

func Getgraph(w http.ResponseWriter, r *http.Request) {
	//セキュリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")    // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// ヘッダの設定．
	r.Header.Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {

		fmt.Printf("%s %s\n", r.URL.Path, r.Method)
		DayTime, err := db.GetDayTime()
		// エラー処理
		if err != nil {
			// ヘッダーに失敗したことを書き込む
			w.WriteHeader(http.StatusServiceUnavailable)
			// ついでに失敗したことをフロントがJSONとして認識できるように書き込む
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			// ログにも書く
			fmt.Println("database error(DayTime)", err)
			// 終了
			return
		}

		// 取得したタスク一覧をByte列に変換
		jsonBytes, err := json.Marshal(DayTime)
		// エラー処理
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("JSON Marshal error(DayTime)", err)
			return
		}

		// 文字列に変換（送信可能になった）
		jsonString := string(jsonBytes)

		// 成功したことをヘッダーに書き込む
		w.WriteHeader(http.StatusOK)

		// 結果を書き込んで終了．
		fmt.Fprintln(w, jsonString)

	}
}
