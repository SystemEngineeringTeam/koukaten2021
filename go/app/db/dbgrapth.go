package db

import "fmt"

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
	// SQL文作成
	DayTimes := DayTime{}
	DayTimes.Monday = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	DayTimes.Tuesday = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	DayTimes.Wednesday = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	DayTimes.Thursday = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	DayTimes.Friday = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	DayTimes.Saturday = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	DayTimes.Sunday = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}

	query := `SELECT AVG(people_count) FROM people WHERE TIME(datetime) BETWEEN '05:00:00' AND '05:59:00' AND DAYOFWEEK(datetime) = '5';`
	//SELECT AVG(people_count) FROM people WHERE TIME(datetime) BETWEEN '05:00:00' AND '05:59:00' AND DAYOFWEEK(datetime) = '5';
	// SQL実行
	rows, err := db.Query(query)
	if err != nil {
		return DayTime{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var people_count float32
		err := rows.Scan(&people_count)
		if err != nil {
			return DayTime{}, err
		}

		fmt.Println(people_count)
	}

	return DayTimes, nil
}
