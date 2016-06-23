package main

import (
	"fmt"
	"testing"
)

func TestGetDisplayName(t *testing.T) {
	t.Skip()
	u := User{FirstName: "Rahul", LastName: "Singh"}

	if u.DisplayName() != "Singh Rahul" {
		t.Errorf("Display name is not in LastName FirstName format, expected %s got %s", "Singh Rahul", u.DisplayName())
	}
}

func TestGetDisplayNameFrenchStyle(t *testing.T) {
	t.Skip()
	u := User{FirstName: "Rahul", LastName: "Singh"}

	if u.DisplayNameFrenchStyle() != "SINGH Rahul" {
		t.Errorf("Display name is not in LASTNAME FirstName format, expected %s got %s", "SINGH Rahul", u.DisplayName())
	}
}

func TestGetCustomDisplayName(t *testing.T) {
	u := User{FirstName: "Rahul", LastName: "Singh"}
	// u.CustomAction = a
	cd, _ := u.GetCustomDisplay()

	if cd != "Rahul-Singh" {
		t.Errorf("Display name is not in FirstName-LastName format, expected %s got %s", "Rahul-Singh", cd)
	}
}

// func a(u User) string {
// 	return fmt.Sprintf("%s-%s", u.FirstName, u.LastName)
// }
