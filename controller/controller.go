package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController    home
	recordsController records
)

func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	recordsController.recordsTemplate = templates["records.html"]
	homeController.loginTemplate = templates["login.html"]
	homeController.registerRoutes()
	recordsController.registerRoutes()
	http.Handle("/style/", http.FileServer(http.Dir(".")))
}
