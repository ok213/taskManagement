package models

import (
	"fmt"
	"log"
	"strings"
)

// User ....
type User struct {
	ID         int
	Email      string
	Hash       string
	Role       string
	FirstName  string
	LastName   string
	MiddleName string
	ShortName  string
	Department string
}

func getShortName(fn, ln, mn string) string {
	f := strings.ToUpper(string([]rune(fn)[0]))
	m := strings.ToUpper(string([]rune(mn)[0]))
	return fmt.Sprintf("%s %s.%s.", ln, f, m)
}

// GetUser ...
func (db *DB) GetUser(email string) (*User, error) {
	var user User

	q := `
	SELECT users.id, users.email, users.hash, roles.role
	FROM users
	LEFT JOIN roles ON users.role = roles.id
	WHERE users.email = $1
	;`

	err := db.QueryRow(q, email).Scan(&user.ID, &user.Email, &user.Hash, &user.Role)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	q = `
	SELECT users.firstname, users.lastname, users.middlename, departments.department
	FROM users
	LEFT JOIN departments ON users.department = departments.id
	WHERE email = $1
	;`

	err = db.QueryRow(q, email).Scan(&user.FirstName, &user.LastName, &user.MiddleName, &user.Department)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	user.ShortName = getShortName(user.FirstName, user.LastName, user.MiddleName)

	return &user, nil
}
