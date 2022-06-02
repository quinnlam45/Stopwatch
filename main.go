package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pluralsight/webservice/controller"
	"github.com/pluralsight/webservice/model"

	"database/sql"

	mssql "github.com/denisenkom/go-mssqldb"
)

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	// results, err := model.GetRecords()
	// if err != nil {
	// 	fmt.Printf("Result error: %v", err)
	// }
	// fmt.Printf("Results: %v", results)
	controller.Startup(templates)
	http.ListenAndServe(":8000", nil) //new(middleware.TimeoutMiddleware)
}

func connectToDatabase() *sql.DB {
	connString := "odbc:server=DESKTOP-TC8EDHO\\SQLEXPRESS;user id=stopwatch_sa;password=stopwatch;database=Stopwatch"

	connector, err := mssql.NewConnector(connString)
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to connect to database: %v", err))
	}
	db := sql.OpenDB(connector)
	model.SetDatabase(db)
	return db
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath + "/_header.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}
