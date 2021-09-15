package db

// Peopleは実際にクライアントに返信するデータの構造体
type People struct {
	Status string `json:"status"`
	People int    `json:"people"`
	Date   string `json:"date"`
}

// GetPeopleはDBから最新のPeopleを取得する関数
func GetPeople() (People, error) {
	rows, err := db.Query("select people_count, datetime from people order by datetime desc limit 1;")
	defer rows.Close()
	if err != nil {
		return People{}, err
	}
	var people People
	rows.Next()
	rows.Scan(&people.People, &people.Date)
	return people, nil
}
