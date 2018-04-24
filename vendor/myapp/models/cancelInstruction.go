package models

import (
	"log"
	"time"
)

// CancelInstruction ...
func (db *DB) CancelInstruction(id int, t time.Time) error {
	q := `
	UPDATE instructions
	SET status = (SELECT statuses.id FROM statuses WHERE statuses.status = 'отменено')
	  , finish = $1
	WHERE instructions.id = $2
	;`

	_, err := db.Exec(q, t, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
