package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
)

// AdminAddDept ...
func (c *AppContext) AdminAddDept(w http.ResponseWriter, r *http.Request) {
	log.Println("AdminAddDept HandleFunc executing ...")

	if r.Method == "POST" {
		r.ParseForm()

		dept := models.Department{
			0,
			r.Form.Get("department"),
		}

		err := c.Db.AddDept(&dept)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/admin/departments", http.StatusSeeOther)
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
