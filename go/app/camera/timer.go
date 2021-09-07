package camera

import (
	"log"
	"time"

	"set1.ie.aitech.ac.jp/koukaten2021/db"
)

// タイムゾーンの定義
var (
	jst = time.FixedZone("Asia/Tokyo", 9*60*60)
)

// カメラ起動処理着火間隔(分)
// 起動間隔は60の約数でなければならない
const IGNITION_INTERVAL = 30

// ResetTimer はリセットを呼び続ける関数
func CameraTimer() {
	t := time.Now().In(jst)        // 現在時刻を取得
	nextDate := getNextDate(t)     // 次の日を取得
	callResetFunc(nextDate.Sub(t)) // リセットし続けるための関数を呼ぶ
}

func getNextDate(t time.Time) time.Time {
	minute := t.Minute()                                                    // 現在の分を取得
	minute = ((int)(minute/IGNITION_INTERVAL)+1)*IGNITION_INTERVAL - minute // 発火させる時間までの分を計算

	nextDatehoge := t.Add(time.Duration(minute) * time.Minute) // minute後のデータを取得
	Y, M, D := nextDatehoge.Date()                             // 年月日を取得
	h := nextDatehoge.Hour()                                   // 時間を取得
	m := nextDatehoge.Minute()                                 // 分を取得
	nextDate := time.Date(Y, M, D, h, m, 0, 0, jst)            // 次の発火時間を取得(秒を切り捨て)
	return nextDate
}

func callResetFunc(d time.Duration) {
	timer := time.NewTimer(d)

	for t := range timer.C { // timer.Cはチャンネルから送られた現在時刻
		// ここでカメラを起動する関数
		people, err := StartCamera()
		if err != nil {
			log.Println("カメラの読み込みに失敗しました")
			log.Println(err)
		} else { // テーブルにinsert
			if db.InsertLog(people, t) != nil {
				log.Println("データベースでの追加に失敗しました")
				log.Println(err)
			}
		}
		timer.Reset(getNextDate(t).Sub(t)) // 次の日の00:00:00に発火するように設定
	}
}
