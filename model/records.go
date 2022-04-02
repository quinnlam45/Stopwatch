package model

import (
	"fmt"
	"time"
)

type Record struct {
	Date     *time.Time
	Time     *time.Time
	Distance string
	By       string
}

func GetRecords() ([]*Record, error) {
	results := []*Record{}
	record_result := &Record{}
	row := db.QueryRow(`EXEC spGetRecords 1`)
	if err := row.Scan(&record_result.Date, &record_result.Time, &record_result.Distance, &record_result.By); err != nil {
		return nil, err
	}
	fmt.Println(record_result)
	results = append(results, record_result)
	fmt.Println(results)
	return results, nil
}
