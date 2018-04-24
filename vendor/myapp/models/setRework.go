package models

import (
	"log"
	"time"
)

// SetRework ...
type SetRework struct {
	IDInstr  int
	Date     time.Time
	Calldown string
}

// SetRework ...
func (db *DB) SetRework(rework SetRework) error {

	q := `
	UPDATE answers SET datecalldown = $1, calldown = $2 WHERE answers.id = (
		SELECT answers.id
		FROM answers
		WHERE answers.idinstr = $3
		ORDER BY answers.date DESC
		LIMIT 1)
	;`

	_, err := db.Exec(q, rework.Date, rework.Calldown, rework.IDInstr)
	if err != nil {
		log.Println(err)
		return err
	}

	q = `
	UPDATE instructions
	SET status = (SELECT statuses.id FROM statuses WHERE statuses.status = 'выполняется доработка')
	WHERE instructions.id = $1
	;`
	_, err = db.Exec(q, rework.IDInstr)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
