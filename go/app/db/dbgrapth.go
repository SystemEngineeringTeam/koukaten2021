package db

import (
	"database/sql"
	"fmt"
)

//構造体の中に配列を格納する

type GraphPeople struct {
	Status    string `json:"status"`
	Monday    []int  `json:"monday"`
	Tuesday   []int  `json:"tuesday"`
	Wednesday []int  `json:"wednesday"`
	Thursday  []int  `json:"thursday"`
	Friday    []int  `json:"friday"`
	Saturday  []int  `json:"saturday"`
	Sunday    []int  `json:"sunday"`
}

//データベースからからSQLを実行し，曜日と時間毎に配列に格納する
func GetDayTime() (GraphPeople, error) {

	GraphPeoples := GraphPeople{}

	// 月曜日から日曜日まで1時間毎に時間を格納する処理
	for j := 1; j < 8; j++ {
		for i := 0; i < 24; i++ {
			// SQL文作成
			Query := fmt.Sprintf("SELECT AVG(people_count) FROM people WHERE HOUR(datetime) = '%d' AND DAYOFWEEK(datetime) = '%d';", i, j)
			// SQL実行
			rows, err := db.Query(Query)
			if err != nil {
				return GraphPeople{}, err
			}

			// SQLを実行した結果を取得
			defer rows.Close()
			var people_count sql.NullFloat64 = sql.NullFloat64{}
			err = rows.Scan(&people_count)
			if err != nil {
				return GraphPeople{}, err
			}

			var people_countValid float64
			if !people_count.Valid {
				people_countValid = 0
			} else {
				people_countValid = people_count.Float64
			}

			switch j {
			case 1:
				GraphPeoples.Monday = append(GraphPeoples.Monday, int(people_countValid))
			case 2:
				GraphPeoples.Tuesday = append(GraphPeoples.Tuesday, int(people_countValid))
			case 3:
				GraphPeoples.Wednesday = append(GraphPeoples.Wednesday, int(people_countValid))
			case 4:
				GraphPeoples.Thursday = append(GraphPeoples.Thursday, int(people_countValid))
			case 5:
				GraphPeoples.Friday = append(GraphPeoples.Friday, int(people_countValid))
			case 6:
				GraphPeoples.Saturday = append(GraphPeoples.Saturday, int(people_countValid))
			case 7:
				GraphPeoples.Sunday = append(GraphPeoples.Sunday, int(people_countValid))
			}
		}
	}
	// 成功した場合において200 OK
	GraphPeoples.Status = "200 OK"
	// 格納したデータを返す
	return GraphPeoples, nil
}
