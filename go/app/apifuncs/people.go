package apifuncs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
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

		var Guests []Guest

		// TODO:関数からexec使える様にする
		out, err := exec.Command("python", "../../python/detect.py").Output()
		if err != nil {
			fmt.Println("Command Exec Error.")
		}
		// 実行したコマンドの結果を出力
		fmt.Printf("\n%s", string(out))

		Guests = append(Guests, Guest{People: 114514})
		//fmt.Print(Guests)

		jsonBytes, err := json.Marshal(Guests)
		//エラー処理
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"Unavailable"}`)
			fmt.Println("JSON Marshal error(Tasks)", err)
			return
		}

		jsonString := string(jsonBytes)

		w.WriteHeader(http.StatusOK)

		if Guests == nil {
			jsonString = "[]"
		}

		fmt.Fprintln(w, jsonString)
	}
}
