package viewmodel

import (
	"github.com/pluralsight/webservice/model"
)

type Records struct {
	Title      string
	Active     string
	RecordList []*Record
}

type Record struct {
	Date         string
	Time         string
	Distance     float32
	DistanceUnit string
	By           string
}

func NewRecords() Records {
	results := Records{
		Title:  "Records",
		Active: "records",
	}
	records, _ := model.GetAllRecords()

	for _, record := range records {
		display_record := &Record{}
		display_record.Date = record.Date.Format("02 Jan 2006")
		display_record.Time = record.Time.Format("15:04:05")
		display_record.Distance = record.Distance
		display_record.DistanceUnit = record.DistanceUnit
		display_record.By = record.By
		results.RecordList = append(results.RecordList, display_record)
	}

	return results
}
