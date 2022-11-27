package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController    home
	recordsController records
	userController    user
)

func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	homeController.loginTemplate = templates["login.html"]
	recordsController.recordsTemplate = templates["records.html"]
	userController.userTemplate = templates["add-user.html"]
	homeController.registerRoutes()
	recordsController.registerRoutes()
	userController.registerRoutes()
	http.Handle("/style/", http.FileServer(http.Dir(".")))
}
