package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
)

type pageIncomingValues struct {
	*MyCustomClaims
	Instructions *[]models.GetInstruction
	Error        string
}

// Incoming выводит информацию по входящим поручениям
func (c *AppContext) Incoming(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming HandleFunc executing ...")

	if r.Method == "GET" {
		key := keyContext("Values")
		claims := r.Context().Value(key).(*MyCustomClaims)

		instructions, err := c.Db.GetIncomeInstructions(claims.ID)
		if err != nil {
			c.Tmpl.ExecuteTemplate(w, "incoming", pageIncomingValues{claims, nil, err.Error()})
			return
		}

		c.Tmpl.ExecuteTemplate(w, "incoming", pageIncomingValues{claims, instructions, ""})
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
