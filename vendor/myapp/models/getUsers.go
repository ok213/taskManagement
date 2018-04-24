package models

import (
	"log"
)

// GetUsers ...
func (db *DB) GetUsers() (*[]User, error) {
	var users []User
	q := `
	SELECT users.id, users.email, roles.role, users.firstname, users.lastname, users.middlename, departments.department
	FROM users
	LEFT JOIN roles ON users.role = roles.id
	LEFT JOIN departments ON users.department = departments.id
	ORDER BY users.lastname ASC
	;`

	rows, err := db.Query(q)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Role, &user.FirstName, &user.LastName, &user.MiddleName, &user.Department); err != nil {
			log.Println(err)
		}

		user.ShortName = getShortName(user.FirstName, user.LastName, user.MiddleName)

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	return &users, nil
}
