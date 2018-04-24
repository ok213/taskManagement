package models

import (
	"log"
)

// User ....
type UserDeptName struct {
	Id         int
	FirstName  string
	LastName   string
	MiddleName string
	Department string
}

// GetUsersExcept ...
func (db *DB) GetUsersExcept(email string) ([]UserDeptName, error) {
	users := make([]UserDeptName, 0)

	q := `
	SELECT users.id, users.lastname, users.firstname, users.middlename, departments.department
	FROM users
	LEFT JOIN departments ON users.department = departments.id
	WHERE email != $1 AND email != 'admin@local.loc'
	ORDER BY users.lastname ASC
	;`

	rows, err := db.Query(q, email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user UserDeptName
		if err := rows.Scan(&user.Id, &user.LastName, &user.FirstName, &user.MiddleName, &user.Department); err != nil {
			log.Fatal(err)
		}

		// users = append(users, fmt.Sprintf("%s %s %s (%s)", user.LastName, user.FirstName, user.MiddleName, user.Department))
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil
}
