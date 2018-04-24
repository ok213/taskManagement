package models

import (
	"log"
	"time"
)

// SetInstruction ...
type SetInstruction struct {
	Owner    int
	Executor int
	Start    time.Time
	Period   time.Time
	Content  string
}

// SetInstruction ...
func (db *DB) SetInstruction(instr SetInstruction) error {

	q := `
	INSERT INTO instructions (pid, owner, executor, start, period, prolonged, finish, status, content)
	VALUES (0, $1, $2, $3, $4, $5, $6, 1, $7)
	;`
	_, err := db.Exec(q, instr.Owner, instr.Executor, instr.Start, instr.Period, instr.Period, instr.Period, instr.Content)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
