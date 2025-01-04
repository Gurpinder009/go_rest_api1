package models

import "strings"

type User struct {
	Fname    string
	Lname    string
	Email    string
	Phone    string
	password string
}

func (user User) GetName() string {
	return user.Fname + " " + user.Lname
}

func (user User) GetEmail() string {
	return user.Email
}

func (user User) GetPhone() string {
	return user.Phone
}

func (user *User) SetName(name string) {
	temp := strings.Split(name, " ")
	user.Fname = temp[0]
	user.Lname = temp[1]
}

func (user *User) SetEmail(Email string) {
	user.Email = Email
}

func (user *User) SetPassword(password string) {
	user.password = password
}

func (user User) MatchPassword(pass string) bool {
	return pass == user.password
}
