package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/pluralsight/webservice/model"
	"github.com/pluralsight/webservice/viewmodel"
)

type user struct {
	userTemplate *template.Template
}

func (u user) registerRoutes() {
	http.HandleFunc("/add-user", u.handleAddUser)
}

func (u user) handleAddUser(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewAddUser()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("error: %v", err))
		}
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		model.AddUser(username, password)
	}
	u.userTemplate.Execute(w, vm)
}
