package models

import (
	"log"
)

// AddDept ...
func (db *DB) AddDept(dept *Department) error {
	q := `
	INSERT INTO departments (department)
	VALUES ($1)
	;`
	_, err := db.Exec(q, dept.Department)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
