package models

import (
	"log"
)

// AddRole ...
func (db *DB) AddRole(role *Role) error {
	q := `
	INSERT INTO roles (role)
	VALUES ($1)
	;`
	_, err := db.Exec(q, role.Role)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
