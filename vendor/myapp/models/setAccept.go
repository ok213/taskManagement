package models

import (
	"log"
	"time"
)

// SetAccept ...
func (db *DB) SetAccept(id int) error {
	t := time.Now()

	q := `
	UPDATE instructions
	SET status = (SELECT statuses.id FROM statuses WHERE statuses.status = 'исполнено')
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
