package models

import (
	"log"
)

// Role ...
type Department struct {
	ID         int
	Department string
}

// GetRole ...
func (db *DB) GetDepartments() (*[]Department, error) {
	var departments []Department
	q := `
	SELECT departments.id, departments.department
	FROM departments
	ORDER BY departments.id ASC
	;`

	rows, err := db.Query(q)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var department Department
		if err := rows.Scan(&department.ID, &department.Department); err != nil {
			log.Println(err)
		}

		departments = append(departments, department)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	return &departments, nil
}
