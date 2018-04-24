package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
)

// AdminAddStatus ...
func (c *AppContext) AdminAddStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("AdminAddStatus HandleFunc executing ...")

	if r.Method == "POST" {
		r.ParseForm()

		status := models.Status{
			0,
			r.Form.Get("status"),
		}

		err := c.Db.AddStatus(&status)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/admin/statuses", http.StatusSeeOther)
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
