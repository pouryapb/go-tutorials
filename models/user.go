package models

import (
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
