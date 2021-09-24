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
func GetGraphPeople() (GraphPeople, error) {
	// 配列の初期化
	GraphPeoples := GraphPeople{}

	// SQLを実行し，その分の数だけ取得
	Query := fmt.Sprintf("SELECT DAYOFWEEK(datetime), AVG(people_count) FROM people GROUP BY DAYOFWEEK(datetime), HOUR(datetime);")
	// SQL実行
	rows, err := db.Query(Query)
	if err != nil {
		return GraphPeople{}, err
	}

	defer rows.Close()

	// SQLで取得した行数分のみを実行
	for rows.Next() {
		// SQLの結果一行分を受け取る変数を宣言
		var dayofweek int
		var people_count sql.NullFloat64 = sql.NullFloat64{}

		// SQLの結果を受け取る
		err = rows.Scan(&dayofweek, &people_count)
		if err != nil {
			return GraphPeople{}, err
		}

		// nullの場合は0を代入
		var people_countValid float64
		if !people_count.Valid {
			people_countValid = 0
		} else {
			people_countValid = people_count.Float64
		}

		// 曜日と時間に対応する配列に格納
		switch dayofweek {
		case 1:
			GraphPeoples.Sunday = append(GraphPeoples.Sunday, int(people_countValid))
		case 2:
			GraphPeoples.Monday = append(GraphPeoples.Monday, int(people_countValid))
		case 3:
			GraphPeoples.Tuesday = append(GraphPeoples.Tuesday, int(people_countValid))
		case 4:
			GraphPeoples.Wednesday = append(GraphPeoples.Wednesday, int(people_countValid))
		case 5:
			GraphPeoples.Thursday = append(GraphPeoples.Thursday, int(people_countValid))
		case 6:
			GraphPeoples.Friday = append(GraphPeoples.Friday, int(people_countValid))
		case 7:
			GraphPeoples.Saturday = append(GraphPeoples.Saturday, int(people_countValid))
		}
	}

	// 成功した場合において"200 OK"と返す
	GraphPeoples.Status = "200 OK"
	// 格納したデータを返す
	return GraphPeoples, nil
}
