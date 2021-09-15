package apifuncs

import (
	"encoding/json"
	"fmt"
	"net/http"

	"set1.ie.aitech.ac.jp/koukaten2021/db"
)

// Getpeopleはクライアントに対して人数と日付を伝える関数
func Getpeople(w http.ResponseWriter, r *http.Request) {
	// アクセスを許可したいアクセス元
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//アクセスを許可したいHTTPメソッド
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// 許可したいHTTPリクエストヘッダ
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// ヘッダの設定
	r.Header.Set("Content-Type", "application/json")

	// GET
	if r.Method == http.MethodGet {
		people, err := db.GetPeople()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"500 INTERNAL SERVER ERROR","message":"Can't get people from database."}`)
			fmt.Println("Can't get people from database.\n", err)
			return
		}

		// 実行したコマンドの結果を出力
		fmt.Printf("\n%d", people.People)
		people.Status = "200 OK"

		jsonBytes, err := json.Marshal(people)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"500 INTERNAL SERVER ERROR","message":"JSON Marshal error(People)"}`)
			fmt.Println("JSON Marshal error(People)\n", err)
			return
		}

		jsonString := string(jsonBytes)

		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, jsonString)
	}
}
