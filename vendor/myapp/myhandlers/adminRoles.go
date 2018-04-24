package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
)

type pageAdminRoles struct {
	*MyCustomClaims
	Roles *[]models.Role
	Error string
}

// AdminRoles ...
func (c *AppContext) AdminRoles(w http.ResponseWriter, r *http.Request) {
	log.Println("AdminRoles Handler executing...")

	if r.Method == "GET" {
		key := keyContext("Values")
		claims := r.Context().Value(key).(*MyCustomClaims)

		roles, err := c.Db.GetRoles()
		if err != nil {
			log.Printf("AdminRoles Handler Error: %v\n", err.Error())
			c.Tmpl.ExecuteTemplate(w, "adminRoles", pageAdminRoles{claims, nil, err.Error()})
			return
		}

		c.Tmpl.ExecuteTemplate(w, "adminRoles", pageAdminRoles{claims, roles, ""})
		return

	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
