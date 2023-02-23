package model

import (
	"testing"
)

func TestLogin(t *testing.T) {
	testUser := "testing_user"
	testPassword := "Test_Password_123"
	result, err := Login(testUser, testPassword)
	// expectedUserResult := &User{Username: testUser, Password: "$2a$12$Ol7CXY6VHpkiSLlcOidfGeefAQjCveM/t382Wrf2.WjQL9xbbDMou"}

	// if err != nil {
	// 	fmt.Println(err)
	// } else if expectedUserResult.Password != passwordCompareResult {
	// 	t.Errorf("Result was incorrect, got: %s, want: %s.", result, expectedHash)
	// }
}
