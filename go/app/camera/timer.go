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

// CameraTimer はカメラを定期的に起動する関数
func CameraTimer() {
	startDateTime := time.Now().In(jst) // カメラ定期起動処理開始時の時刻を取得

	timer := time.NewTimer(getNextTime(startDateTime).Sub(startDateTime)) // 次の起動時刻にタイマーをセット

	for t := range timer.C { // timer.Cはチャンネルから送られた現在時刻
		log.Println("カメラ起動")

		// 撮影時間による誤差の吸収のため，カメラの起動より先に次の起動時刻にタイマーをセット．
		timer.Reset(getNextTime(t).Sub(t))

		// カメラを起動する関数
		people, err := StartCamera()
		if err != nil { // カメラの起動に失敗した場合
			// 失敗を報告しつつも処理は続行する
			log.Println("カメラの起動に失敗しました")
			log.Println(err)
		} else { // テーブルにinsert
			if db.InsertLog(people, t) != nil { // insertに失敗した場合
				// 失敗を報告しつつも処理は続行する
				log.Println("データベースでの追加に失敗しました")
				log.Println(err)
			}
		}
	}
}

// getNextTime は次のカメラ起動時刻を取得する関数
func getNextTime(t time.Time) time.Time {
	minute := t.Minute()                                                    // 現在の分を取得
	minute = ((int)(minute/IGNITION_INTERVAL)+1)*IGNITION_INTERVAL - minute // 次の起動時刻までの分を計算

	// 年月日,時間,分のみ取得し，秒,ミリ秒は切り捨てる．
	nextDate := t.Add(time.Duration(minute) * time.Minute) // minute後の時刻を取得
	Y, M, D := nextDate.Date()                             // 年月日を取得
	h := nextDate.Hour()                                   // 時間を取得
	m := nextDate.Minute()                                 // 分を取得

	// 次の起動時刻を取得(秒を切り捨て)
	// ここで，秒を1にすることによって，確実に起動時刻に乗ってからカメラを起動するように調節している．
	nextDate = time.Date(Y, M, D, h, m, 1, 0, jst)
	return nextDate
}
