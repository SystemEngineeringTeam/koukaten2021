package db

import (
	"fmt"
	"math"
)

//構造体の中に配列を格納する

type DayTime struct {
	Monday    []int `json:"monday"`
	Tuesday   []int `json:"tuesday"`
	Wednesday []int `json:"wednesday"`
	Thursday  []int `json:"thursday"`
	Friday    []int `json:"friday"`
	Saturday  []int `json:"saturday"`
	Sunday    []int `json:"sunday"`
}

//データベースからからSQLを実行し，曜日と時間毎に配列に格納する
func GetDayTime() (DayTime, error) {

	DayTimes := DayTime{}
	//　構造体を24時間分作成する
	for i := 0; i < 24; i++ {
		DayTimes.Monday = append(DayTimes.Monday, 0)
		DayTimes.Tuesday = append(DayTimes.Tuesday, 0)
		DayTimes.Wednesday = append(DayTimes.Wednesday, 0)
		DayTimes.Thursday = append(DayTimes.Thursday, 0)
		DayTimes.Friday = append(DayTimes.Friday, 0)
		DayTimes.Saturday = append(DayTimes.Saturday, 0)
		DayTimes.Sunday = append(DayTimes.Sunday, 0)
	}

	// 月曜日から日曜日まで1時間毎に時間を格納する処理
	for j := 1; j < 8; j++ {
		for i := 0; i < 24; i++ {
			// SQL文作成
			sql := fmt.Sprintf("SELECT AVG(people_count) FROM people WHERE HOUR(datetime) = '%d' AND DAYOFWEEK(datetime) = '%d';", i, j)
			// SQL実行
			rows, err := db.Query(sql)
			if err != nil {
				switch j {
				case 1:
					DayTimes.Monday[i] = 0
				case 2:
					DayTimes.Tuesday[i] = 0
				case 3:
					DayTimes.Wednesday[i] = 0
				case 4:
					DayTimes.Thursday[i] = 0
				case 5:
					DayTimes.Friday[i] = 0
				case 6:
					DayTimes.Saturday[i] = 0
				case 7:
					DayTimes.Sunday[i] = 0
				default:
					return DayTime{}, err
				}
				fmt.Printf("エラーが発生しました。 %d曜日の%d時に0を格納します。", j, i)
			}
			// SQLを閉じる
			defer rows.Close()

			// SQLで取得した行数だけ実行
			for rows.Next() {
				var people_count float64
				err := rows.Scan(&people_count)
				if err != nil {
					switch j {
					case 1:
						DayTimes.Monday[i] = 0
					case 2:
						DayTimes.Tuesday[i] = 0
					case 3:
						DayTimes.Wednesday[i] = 0
					case 4:
						DayTimes.Thursday[i] = 0
					case 5:
						DayTimes.Friday[i] = 0
					case 6:
						DayTimes.Saturday[i] = 0
					case 7:
						DayTimes.Sunday[i] = 0
					default:
						return DayTime{}, err
					}
					fmt.Printf("エラーが発生しました。 %d曜日の%d時に0を格納します。", j, i)
				}
				//四捨五入
				people_count = math.Round(people_count)
				// 時間と曜日を比較し，その時間に格納する
				// case文
				switch j {
				case 1:
					DayTimes.Monday[i] = int(people_count)
				case 2:
					DayTimes.Tuesday[i] = int(people_count)
				case 3:
					DayTimes.Wednesday[i] = int(people_count)
				case 4:
					DayTimes.Thursday[i] = int(people_count)
				case 5:
					DayTimes.Friday[i] = int(people_count)
				case 6:
					DayTimes.Saturday[i] = int(people_count)
				case 7:
					DayTimes.Sunday[i] = int(people_count)
				default:
					return DayTime{}, err
				}
			}
		}
	}
	// 格納したデータを返す
	return DayTimes, nil
}
