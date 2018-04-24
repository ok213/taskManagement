package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
)

// AdminAddRole ...
func (c *AppContext) AdminAddRole(w http.ResponseWriter, r *http.Request) {
	log.Println("AdminAddRole HandleFunc executing ...")

	if r.Method == "POST" {
		r.ParseForm()

		role := models.Role{
			0,
			r.Form.Get("role"),
		}

		err := c.Db.AddRole(&role)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/admin/roles", http.StatusSeeOther)
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
