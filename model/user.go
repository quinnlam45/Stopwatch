package model

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    int
	Username  string
	Password  string
	LastLogin *time.Time
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash), err
}

func UpdateLastLogin(username, loginTime string) error {
	_, err := db.Exec(`
	UPDATE [user]
	SET lastLogin = @p1
	WHERE username = @p2`, loginTime, username)

	return err
}

func Login(username, password string) (*User, error) {
	result := &User{}
	var hashedPassword string

	row := db.QueryRow(`
	SELECT userID, username, password
	FROM [user]
	WHERE username = @p1`, username)

	fmt.Printf("%v", row)

	err := row.Scan(&result.UserID, &result.Username, &hashedPassword)

	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("user not found")
	case err != nil:
		return nil, err
	}

	matchResult := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if matchResult == nil {
		fmt.Println("Match!")

		loginTime := time.Now().Format("2006-01-02 15:04:05")
		UpdateLastLogin(result.Username, loginTime)
	} else {
		fmt.Printf("Passwords did not match %v", matchResult)
	}

	return result, nil
}

func AddUser(username, password string) {
	pwdHash, err := HashPassword(password)

	if err != nil {
		fmt.Println(err)
	} else {
		_, err := db.Exec(`EXEC spAddUser @p1, @p2`, username, pwdHash)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("User: %v created successfully!", username)
	}
}
