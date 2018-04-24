package models

import (
	"log"
)

// Role ...
type Status struct {
	ID     int
	Status string
}

// GetStatuses ...
func (db *DB) GetStatuses() (*[]Status, error) {
	var statuses []Status
	q := `
	SELECT statuses.id, statuses.status
	FROM statuses
	ORDER BY statuses.id ASC
	;`

	rows, err := db.Query(q)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var status Status
		if err := rows.Scan(&status.ID, &status.Status); err != nil {
			log.Println(err)
		}

		statuses = append(statuses, status)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	return &statuses, nil
}
