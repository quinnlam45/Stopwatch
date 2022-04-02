package viewmodel

type Records struct {
	Title      string
	Active     string
	RecordList []Record
}

type Record struct {
	Date     string
	Time     string
	Distance string
	By       string
}

func NewRecords() Records {
	results := Records{
		Title:  "Records",
		Active: "records",
	}
	record1 := Record{
		Date:     "31/1/2022",
		Time:     "15:20.12",
		Distance: "4.5mi",
		By:       "Run",
	}
	record2 := Record{
		Date:     "19/2/2022",
		Time:     "12:24.00",
		Distance: "3.0mi",
		By:       "Bike",
	}

	results.RecordList = []Record{record1, record2}

	return results
}

// func NewRecords() []*model.Record {
// 	record, _ := model.GetRecords()
// 	return record
// }
