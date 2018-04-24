package models

import (
	"log"
)

// AddStatus ...
func (db *DB) AddStatus(status *Status) error {
	q := `
	INSERT INTO statuses (status)
	VALUES ($1)
	;`
	_, err := db.Exec(q, status.Status)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
