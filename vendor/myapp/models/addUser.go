package models

import (
	"log"
)

// AddUser ...
func (db *DB) AddUser(user *User) error {
	q := `
	INSERT INTO users (email, hash, role, firstname, lastname, middlename, department)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	;`
	_, err := db.Exec(q, user.Email, user.Hash, user.Role, user.FirstName, user.LastName, user.MiddleName, user.Department)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
