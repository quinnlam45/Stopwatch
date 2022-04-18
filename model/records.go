package model

import (
	"fmt"
	"time"
)

type Record struct {
	Date         *time.Time
	Time         *time.Time
	Distance     float32
	DistanceUnit string
	By           string
}

func GetRecord() (*Record, error) {
	record_result := &Record{}
	row := db.QueryRow(`EXEC spGetRecord 1`)
	if err := row.Scan(&record_result.Date, &record_result.Time, &record_result.Distance, &record_result.By); err != nil {
		fmt.Println(err)
		return nil, err
	}
	//fmt.Println(record_result)
	return record_result, nil
}

func GetAllRecords() ([]*Record, error) {
	rows, err := db.Query(`EXEC spGetAllRecords`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	records := []*Record{}
	for rows.Next() {
		record_result := &Record{}
		err := rows.Scan(&record_result.Date, &record_result.Time, &record_result.Distance, &record_result.DistanceUnit, &record_result.By)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		records = append(records, record_result)
	}
	return records, nil
}
