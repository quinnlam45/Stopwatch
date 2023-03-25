package model

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	testDB := new(mockDB)
	db = testDB

	testDB.setMockDBRow(5, "test_user", "somehash")
	userMockRow := testDB.returnedRow

	username := "test_user"
	test_password := "Password"
	result, err := Login(username, test_password)
	if err != nil {
		fmt.Println(err)
	}

	if result.UserID != userMockRow.values[0] && result.Username != userMockRow.values[0] {
		t.Errorf(
			`Want id: 5 username: test_user,
				got id: %v username: %v`,
			result.UserID,
			result.Username)
	}

}
