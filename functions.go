package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

}

type User struct {
	FirstName, LastName string
	CustomAction        func(User) string
}

func (u *User) GetCustomDisplay() (string, error) {
	if u.CustomAction != nil {
		return u.CustomAction(*u), nil

	}
	return "", errors.New("No custom action defined")
}

func (u *User) DisplayName() string {
	return fmt.Sprintf("%s %s", u.LastName, u.FirstName)
}

func (u *User) DisplayNameFrenchStyle() string {
	return fmt.Sprintf("%s %s", strings.ToUpper(u.LastName), u.FirstName)
}
