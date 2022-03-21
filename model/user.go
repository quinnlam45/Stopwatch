package model

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"
)

const passwordSalt = "F0461966-50EC-42BB-8370-A0B96A59E533"

type User struct {
	UserID    int
	Username  string
	Password  string
	LastLogin *time.Time
}

func Login(username, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hasher.Write([]byte(passwordSalt))
	// pwd := hex.EncodeToString(hasher.Sum(nil))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	fmt.Printf("%v\n", pwd)

	row := db.QueryRow(`
			SELECT userID, username
			FROM [user]
			WHERE username = $1
				AND password = CAST($2 AS VARBINARY)`, username, pwd)
	fmt.Printf("%v", row)
	err := row.Scan(&result.UserID, &result.Username)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("user not found")
	case err != nil:
		return nil, err
	}
	return result, nil
}
