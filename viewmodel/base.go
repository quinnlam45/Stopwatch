package viewmodel

type Base struct {
	Title  string
	Active string
}

func NewBase() Base {
	return Base{
		Title:  "Stopwatch",
		Active: "stopwatch",
	}
}
