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
	http.Handle("/static/style/", http.StripPrefix("/static/style", http.FileServer(http.Dir("static/style"))))
	http.HandleFunc("/static/script/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/javascript")

		fileserver := http.StripPrefix("/static/script", http.FileServer(http.Dir("static/script")))
		fileserver.ServeHTTP(w, r)
	})
}
