package apifuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//MySQLの事も考えて構造体にあえて実装
type Guest struct {
	People int `json:"people"`
}

//Getで返す関数
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
		out, err := http.Get("http://host.docker.internal:8082/")
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("Can't run detect.py\n", err)
			return
		}
		defer out.Body.Close()

		jsonBytes, err := ioutil.ReadAll(out.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("Can't catch people(io error)\n", err)
			return
		}

		var guest Guest

		if err := json.Unmarshal(jsonBytes, &guest); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("Can't catch people(JSON Unmarshal error)\n", err)
			return
		}

		// 実行したコマンドの結果を出力
		fmt.Printf("\n%d", guest.People)

		jsonBytes, err = json.Marshal(guest)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("JSON Marshal error(People)\n", err)
			return
		}

		jsonString := string(jsonBytes)

		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, jsonString)
	}
}
