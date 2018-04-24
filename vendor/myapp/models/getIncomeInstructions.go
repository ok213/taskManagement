package models

import (
	"log"
	"time"
)

// GetIncomeInstructions ...
func (db *DB) GetIncomeInstructions(id int) (*[]GetInstruction, error) {
	instructions := make([]GetInstruction, 0)
	q := `
	SELECT instructions.id, instructions.period, statuses.status, users.lastname, users.firstname, users.middlename, departments.department, instructions.content
	FROM instructions
	LEFT JOIN users ON instructions.owner = users.id
	LEFT JOIN statuses ON instructions.status = statuses.id
	LEFT JOIN departments ON users.department = departments.id
	WHERE instructions.executor = $1 AND instructions.status != (SELECT statuses.id FROM statuses WHERE statuses.status = 'ожидание принятия исполнения')
									 AND instructions.status != (SELECT statuses.id FROM statuses WHERE statuses.status = 'исполнено')
									 AND instructions.status != (SELECT statuses.id FROM statuses WHERE statuses.status = 'отменено')
	ORDER BY instructions.period ASC;
	`
	rows, err := db.Query(q, id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var instr GetInstruction
		var pd time.Time
		if err := rows.Scan(&instr.ID, &pd, &instr.Status, &instr.LastName, &instr.FirstName, &instr.MiddleName, &instr.Department, &instr.Content); err != nil {
			log.Println(err)
		}

		instr.Period = getDateString(pd)
		instructions = append(instructions, instr)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	return &instructions, nil
}
