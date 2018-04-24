package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
)

type pageAdminStatuses struct {
	*MyCustomClaims
	Statuses *[]models.Status
	Error    string
}

// AdminStatuses ...
func (c *AppContext) AdminStatuses(w http.ResponseWriter, r *http.Request) {
	log.Println("AdminStatuses Handler executing...")

	if r.Method == "GET" {
		key := keyContext("Values")
		claims := r.Context().Value(key).(*MyCustomClaims)

		statuses, err := c.Db.GetStatuses()
		if err != nil {
			log.Printf("AdminStatuses Handler Error: %v\n", err.Error())
			c.Tmpl.ExecuteTemplate(w, "adminStatuses", pageAdminStatuses{claims, nil, err.Error()})
			return
		}

		c.Tmpl.ExecuteTemplate(w, "adminStatuses", pageAdminStatuses{claims, statuses, ""})
		return

	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
