package db

import "time"

const layout = "2006-01-02 15:04:05"

func InsertLog(people int, t time.Time) error {
	_, err := db.Exec("insert into people(datetime, people_count) values(?,?);", t.Format(layout), people)
	return err
}
