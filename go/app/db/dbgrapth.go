package db

type DayTime struct {
	Monday    int `json:"monday"`
	Tuesday   int `json:"tuesday"`
	Wednesday int `json:"wednesday"`
	Thursday  int `json:"thursday"`
	Friday    int `json:"friday"`
	Saturday  int `json:"saturday"`
	Sunday    int `json:"sunday"`
}

//データベースからからSQLを実行し，曜日と時間毎に配列に格納する
func GetDayTime() ([]DayTime, error) {
	// SQL文作成
	var dayTime []DayTime
	query := `SELECT DISTINCT(date), time FROM log`
	// SQL実行
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// SQL結果を配列に格納
	for rows.Next() {
		var Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday int
		err := rows.Scan(&Monday, &Tuesday, &Wednesday, &Thursday, &Friday, &Saturday, &Sunday)
		if err != nil {
			return nil, err
		}
		dayTime = append(dayTime, DayTime{date, time})
	}
	return dayTime, nil
}
