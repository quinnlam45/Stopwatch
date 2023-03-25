package controller

import (
	"html/template"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHomeExecutesCorrectTemplate(t *testing.T) {
	h := new(home)
	expected := `home template`
	h.homeTemplate, _ = template.New("").Parse(expected)
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	h.handleHome(w, r)

	actual, _ := ioutil.ReadAll(w.Result().Body)

	if string(actual) != expected {
		t.Errorf("Failed to execute correct home template")
	}
}

func TestHomePostMethod(t *testing.T) {
	h := new(home)
	r := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()

	form := url.Values{}
	form.Add("form-date", "2023-03-13")
	form.Add("form-time", "10:30:00")
	form.Add("time-record", "01:30:00")
	form.Add("distance", "5.6")
	form.Add("distance-unit", "1")
	form.Add("completed-by", "1")

	r.PostForm = form
	h.handleHome(w, r)

	t.Logf("%v", r.PostForm)

}
