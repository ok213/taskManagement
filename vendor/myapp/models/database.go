package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// Info contains the database configurations
type Info struct {
	// Database type
	Type string
	// PostgreSQL info if used
	PostgreSQL PostgreSQLInfo
}

// PostgreSQLInfo is the details for the database connection
type PostgreSQLInfo struct {
	Hostname   string
	Port       int
	Parameters string
	Dbname     string
	Dbuser     string
	Dbpassword string
}

// DSN returns the string for connect to database
func DSN(ps PostgreSQLInfo) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s %s",
		ps.Hostname, ps.Port, ps.Dbuser, ps.Dbpassword, ps.Dbname, ps.Parameters)
}

// Datastore интерфейс к БД
type Datastore interface {
	// здесь будут добавляться запросы к БД
	GetUser(string) (*User, error)
	GetUsersExcept(string) ([]UserDeptName, error)
	SetInstruction(SetInstruction) error
	GetInstructions(int) (*[]GetInstruction, error)
	GetInstruction(int) (*Instruction, error)
	GetIncomeInstructions(int) (*[]GetInstruction, error)
	GetIncomingInstruction(int) (*Instruction, error)
	SetAnswer(SetAnswer) error
	SetAccept(int) error
	SetRework(SetRework) error
	CancelInstruction(int, time.Time) error
	GetUsers() (*[]User, error)
	GetRoles() (*[]Role, error)
	GetDepartments() (*[]Department, error)
	AddUser(*User) error
	AddDept(*Department) error
	AddRole(*Role) error
	AddStatus(*Status) error
	GetStatuses() (*[]Status, error)
}

// DB ...
type DB struct {
	*sql.DB
}

// Connect to database
func Connect(dbinfo Info) (*DB, error) {
	var db *sql.DB
	var err error
	switch dbinfo.Type {
	case "PostgreSQL":
		db, err = sql.Open("postgres", DSN(dbinfo.PostgreSQL))
		if err != nil {
			log.Println("SQL Driver Error", err)
			return nil, err
		}
		err = db.Ping()
		if err != nil {
			log.Println("Database Error", err)
			return nil, err
		}
		return &DB{db}, nil
	default:
		log.Println("No registered database in config")
		return nil, errors.New("No registered database in config")
	}

}
