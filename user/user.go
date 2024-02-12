package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("invalid user data")
	}

	return &User{
		FirstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
