package models

import (
	"log"
	"time"
)

// GetIncomingInstruction ...
func (db *DB) GetIncomingInstruction(id int) (*Instruction, error) {
	var i Instruction
	q := `
	SELECT instructions.id, instructions.pid, users.lastname, users.firstname, users.middlename, departments.department, instructions.start, instructions.period, instructions.prolonged, statuses.status, instructions.content
	FROM instructions
	LEFT JOIN users ON instructions.owner = users.id
	LEFT JOIN statuses ON instructions.status = statuses.id
	LEFT JOIN departments ON users.department = departments.id
	WHERE instructions.id = $1;
	`
	var st time.Time
	var pd time.Time
	var pl time.Time

	err := db.QueryRow(q, id).Scan(&i.ID, &i.PID, &i.LastName, &i.FirstName, &i.MiddleName, &i.Department, &st, &pd, &pl, &i.Status, &i.Content)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	i.Start = getDateString(st)
	i.Period = getDateString(pd)
	i.Prolonged = getDateString(pl)

	// готовим массив ответов
	answers := make([]GetAnswer, 0)
	q = `
	SELECT answers.date, answers.answer, answers.datecalldown, answers.calldown
	FROM answers
	WHERE answers.idinstr = $1
	ORDER BY answers.date DESC
	;`

	rows, err1 := db.Query(q, id)
	if err1 != nil {
		log.Println(err1)
	}
	defer rows.Close()
	for rows.Next() {
		var answer GetAnswer
		if err := rows.Scan(&st, &answer.Answer, &pd, &answer.Calldown); err != nil {
			log.Println(err)
		}

		answer.Date = getDateString(st)
		answer.DateCalldown = getDateString(pd)
		answers = append(answers, answer)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	i.Answers = &answers

	return &i, nil
}
