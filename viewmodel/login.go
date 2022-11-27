package viewmodel

type Login struct {
	Title    string
	Active   string
	Username string
	Password string
}

func NewLogin() Login {
	return Login{
		Title:  "Login",
		Active: "login",
	}
}

func NewAddUser() Login {
	return Login{
		Title:  "Add user",
		Active: "add user",
	}
}
