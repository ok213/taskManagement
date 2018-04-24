package models

import (
	"log"
	"time"
)

// SetAnswer ...
type SetAnswer struct {
	IDInstr int
	Date    time.Time
	Answer  string
}

// SetAnswer ...
func (db *DB) SetAnswer(answer SetAnswer) error {

	q := `
	INSERT INTO answers (idinstr, date, answer, datecalldown, calldown)
	VALUES ($1, $2, $3, '0001-01-01', '')
	;`
	_, err := db.Exec(q, answer.IDInstr, answer.Date, answer.Answer)
	if err != nil {
		log.Println(err)
		return err
	}

	q = `
	UPDATE instructions
	SET status = (SELECT statuses.id FROM statuses WHERE statuses.status = 'ожидание принятия исполнения')
	WHERE instructions.id = $1
	;`

	_, err = db.Exec(q, answer.IDInstr)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
