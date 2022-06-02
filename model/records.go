package model

import (
	"fmt"
	"time"
)

type Record struct {
	Date           *time.Time
	Time           *time.Time
	Distance       float32
	DistanceUnit   string
	CompletedBy    string
	DistanceUnitID int
	CompletedByID  int
}

func GetRecord() (*Record, error) {
	record_result := &Record{}
	row := db.QueryRow(`EXEC spGetRecord 1`)
	if err := row.Scan(&record_result.Date, &record_result.Time, &record_result.Distance, &record_result.CompletedBy); err != nil {
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
		err := rows.Scan(&record_result.Date, &record_result.Time, &record_result.Distance, &record_result.DistanceUnit, &record_result.CompletedBy)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		records = append(records, record_result)
	}
	return records, nil
}

func AddRecord(formDate time.Time, timeElapsed string, distance float32, distanceUnit int, completedBy int) error {
	//exec spAddRecord '18 Apr 2022', '00:10:00', 1.5, 1, 1
	_, err := db.Query(`EXEC spAddRecord @p1, @p2, @p3, @p4, @p5`, formDate, timeElapsed, distance, distanceUnit, completedBy)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
