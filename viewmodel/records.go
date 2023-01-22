package viewmodel

import (
	"sort"

	"github.com/pluralsight/webservice/model"
)

type Records struct {
	Title      string
	Active     string
	RecordList []*Record
}

type Record struct {
	Date         string  `json:"date"`
	Time         string  `json:"time"`
	Distance     float32 `json:"distance"`
	DistanceUnit string  `json:"distanceUnit"`
	By           string  `json:"by"`
}

func GetRecords() Records {
	results := Records{
		Title:  "Records",
		Active: "records",
	}
	records, _ := model.GetAllRecords()

	for _, record := range records {
		display_record := &Record{}
		display_record.Date = record.Date.Format("02/01/2006")
		display_record.Time = record.Time.Format("15:04:05")
		display_record.Distance = record.Distance
		display_record.DistanceUnit = record.DistanceUnit
		display_record.By = record.CompletedBy
		results.RecordList = append(results.RecordList, display_record)
	}

	return results
}

func FilterRecords(startRange, endRange, ord, colName string) Records {
	results := Records{}
	records, _ := model.GetRecords(startRange, endRange)

	if ord == "desc" {
		switch {
		case colName == "Date":
			sort.SliceStable(records, func(i, j int) bool { comparison := records[j].Date; return records[i].Date.After(*comparison) })
		case colName == "Time":
			sort.SliceStable(records, func(i, j int) bool { comparison := records[j].Time; return records[i].Time.After(*comparison) })
		case colName == "Distance":
			sort.SliceStable(records, func(i, j int) bool { return records[i].Distance > records[j].Distance })
		}
	} else if ord == "asc" {
		switch {
		case colName == "Date":
			sort.SliceStable(records, func(i, j int) bool { comparison := records[j].Date; return records[i].Date.Before(*comparison) })
		case colName == "Time":
			sort.SliceStable(records, func(i, j int) bool { comparison := records[j].Time; return records[i].Time.Before(*comparison) })
		case colName == "Distance":
			sort.SliceStable(records, func(i, j int) bool { return records[i].Distance < records[j].Distance })
		}
	}

	for _, record := range records {
		display_record := &Record{}
		display_record.Date = record.Date.Format("02/01/2006")
		display_record.Time = record.Time.Format("15:04:05")
		display_record.Distance = record.Distance
		display_record.DistanceUnit = record.DistanceUnit
		display_record.By = record.CompletedBy
		results.RecordList = append(results.RecordList, display_record)
		//fmt.Println(display_record)
	}

	return results
}
