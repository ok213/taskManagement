package models

import (
	"log"
)

// Role ...
type Role struct {
	ID   int
	Role string
}

// GetRoles ...
func (db *DB) GetRoles() (*[]Role, error) {
	var roles []Role
	q := `
	SELECT roles.id, roles.role
	FROM roles
	ORDER BY roles.id ASC
	;`

	rows, err := db.Query(q)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var role Role
		if err := rows.Scan(&role.ID, &role.Role); err != nil {
			log.Println(err)
		}

		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	return &roles, nil
}
