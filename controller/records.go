package controller

import (
	"html/template"
	"net/http"

	"github.com/pluralsight/webservice/viewmodel"
)

type records struct {
	recordsTemplate *template.Template
}

func (rec records) registerRoutes() {
	http.HandleFunc("/records", rec.handleRecords)
}

func (rec records) handleRecords(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewRecords()
	rec.recordsTemplate.Execute(w, vm)
}

// func (rec records) handleRecords(w http.ResponseWriter, r *http.Request) {
// 	vm, err := model.GetRecords()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	rec.recordsTemplate.Execute(w, vm)
// }
