package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
	"strconv"
)

type pageIncomingInstr struct {
	*MyCustomClaims
	Instruction *models.Instruction
	Error       string
}

// IncomingInstr выводит информацию по входящим поручениям
func (c *AppContext) IncomingInstr(w http.ResponseWriter, r *http.Request) {
	log.Println("IncomingInstr HandleFunc executing ...")

	if r.Method == "GET" {
		key := keyContext("Values")
		claims := r.Context().Value(key).(*MyCustomClaims)

		ids := r.URL.Query()
		if len(ids) < 1 {
			log.Println("Url Param 'key' is missing")
			http.Redirect(w, r, "/outgoing", http.StatusSeeOther)
			return
		}
		id, _ := strconv.Atoi(ids.Get("ID"))

		instr, err := c.Db.GetIncomingInstruction(id)
		if err != nil {
			c.Tmpl.ExecuteTemplate(w, "incomingInstr", pageIncomingInstr{claims, nil, err.Error()})
			return
		}
		log.Printf("---> instr.answers :  %v \n", instr.Answers)
		c.Tmpl.ExecuteTemplate(w, "incomingInstr", pageIncomingInstr{claims, instr, ""})
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
