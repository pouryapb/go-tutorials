package models

import (
	"errors"

	"github.com/pouryapb/go-tutorials/rest-api/db"
	"github.com/pouryapb/go-tutorials/rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.Database.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashPassword)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	u.Id = id
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"

	row := db.Database.QueryRow(query, u.Email)

	var retrivedPass string
	err := row.Scan(&retrivedPass)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPass)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
