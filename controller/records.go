package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/pluralsight/webservice/viewmodel"
)

type records struct {
	recordsTemplate *template.Template
}

func (rec records) registerRoutes() {
	http.HandleFunc("/records", rec.handleRecords)
	http.HandleFunc("/records/get-all", getRecords)
	http.HandleFunc("/records/filter", filterRecords)
}

func (rec records) handleRecords(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.GetRecords()
	rec.recordsTemplate.Execute(w, vm)
}

func filterRecords(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	startRange := url.Query().Get("from")
	endRange := url.Query().Get("to")
	ord := url.Query().Get("ord")
	colName := url.Query().Get("col")

	records := viewmodel.FilterRecords(startRange, endRange, ord, colName)

	data, err := json.Marshal(records)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getRecords(w http.ResponseWriter, r *http.Request) {
	//url := r.URL
	//ord := url.Query().Get("ord")

	records := viewmodel.GetRecords()

	data, err := json.Marshal(records)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
