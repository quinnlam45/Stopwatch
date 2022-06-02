package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/pluralsight/webservice/model"
	"github.com/pluralsight/webservice/viewmodel"
)

type home struct {
	homeTemplate  *template.Template
	loginTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewBase()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("error: %v", err))
		}
		date, _ := time.Parse("2006-01-02", r.Form.Get("form-date"))
		timeElasped, _ := time.Parse("15:04:05", r.Form.Get("ticker-time"))
		distance, _ := strconv.ParseFloat(r.Form.Get("distance"), 32)
		distanceUnit, _ := strconv.Atoi(r.Form.Get("distance-unit"))
		completedBy, _ := strconv.Atoi(r.Form.Get("completed-by"))

		model.AddRecord(date, timeElasped.Format("15:04:05.000"), float32(distance), distanceUnit, completedBy)
		fmt.Printf("Date %v, Time %v, distance %v, unit %v, by %v", date, timeElasped, float32(distance), distanceUnit, completedBy)
	}
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("error: %v", err))
		}
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if user, err := model.Login(username, password); err == nil {
			log.Printf("user logged in: %v\n", user)
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			log.Printf("failed to log user with username:%v\n error: %v\n", username, err)
			vm.Username = username
			vm.Password = password
		}
	}
	h.loginTemplate.Execute(w, vm)
}
