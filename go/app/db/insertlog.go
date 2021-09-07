package db

import "time"

const layout = "2021-08-15 00:00:00"

func InsertLog(people int, t time.Time) error {
	_, err := db.Exec("insert into people(datatime, people_count) values(?,?);", t.Format(layout), people)
	return err
}
