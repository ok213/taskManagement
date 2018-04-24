package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
)

type pageOutgoingValues struct {
	*MyCustomClaims
	Instructions *[]models.GetInstruction
	Error        string
}

// Outgoing выводит информацию по входящим поручениям
func (c *AppContext) Outgoing(w http.ResponseWriter, r *http.Request) {
	log.Println("Outgoing HandleFunc executing ...")

	if r.Method == "GET" {
		key := keyContext("Values")
		claims := r.Context().Value(key).(*MyCustomClaims)

		instructions, err := c.Db.GetInstructions(claims.ID)
		if err != nil {
			c.Tmpl.ExecuteTemplate(w, "outgoing", pageOutgoingValues{claims, nil, err.Error()})
			return
		}

		c.Tmpl.ExecuteTemplate(w, "outgoing", pageOutgoingValues{claims, instructions, ""})
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
