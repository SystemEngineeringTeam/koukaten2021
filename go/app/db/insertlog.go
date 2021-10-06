package db

func InsertLog(people int, datetime string) error {
	_, err := db.Exec("insert into people(datetime, people_count) values(?,?);", datetime, people)
	return err
}
