package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
	"strconv"
)

type pageOutgoingInstr struct {
	*MyCustomClaims
	Instruction *models.Instruction
	Error       string
}

// OutgoingInstr выводит информацию по входящим поручениям
func (c *AppContext) OutgoingInstr(w http.ResponseWriter, r *http.Request) {
	log.Println("OutgoingInstr HandleFunc executing ...")

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

		instr, err := c.Db.GetInstruction(id)
		if err != nil {
			c.Tmpl.ExecuteTemplate(w, "outgoingInstr", pageOutgoingInstr{claims, nil, err.Error()})
			return
		}

		c.Tmpl.ExecuteTemplate(w, "outgoingInstr", pageOutgoingInstr{claims, instr, ""})
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
