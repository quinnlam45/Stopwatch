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
	_, err := db.Query(`EXEC spAddRecord @p1, @p2, @p3, @p4, @p5`, formDate, timeElapsed, distance, distanceUnit, completedBy)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GetRecords(startRange, endRange string) ([]*Record, error) {
	rows, err := db.Query(`EXEC spGetRecords @p1, @p2`, startRange, endRange)

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
