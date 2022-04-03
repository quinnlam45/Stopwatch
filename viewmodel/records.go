package viewmodel

import (
	"github.com/pluralsight/webservice/model"
)

type Records struct {
	Title      string
	Active     string
	RecordList []*model.Record
}

// type Record struct {
// 	Date     *time.Time
// 	Time     *time.Time
// 	Distance string
// 	By       string
// }

func NewRecords() Records {
	results := Records{
		Title:  "Records",
		Active: "records",
	}
	results.RecordList, _ = model.GetAllRecords()
	// for _, record := range records {
	// 	results.RecordList = append(results.RecordList, record)
	// }

	return results
}

// func NewRecords() []*model.Record {
// 	record, _ := model.GetRecords()
// 	return record
// }
