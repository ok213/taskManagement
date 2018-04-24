package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
)

// AdminAddUser ...
func (c *AppContext) AdminAddUser(w http.ResponseWriter, r *http.Request) {
	log.Println("AdminAddUser HandleFunc executing ...")

	if r.Method == "POST" {
		r.ParseForm()

		user := models.User{
			0,
			r.Form.Get("email"),
			r.Form.Get("password"),
			r.Form.Get("role"),
			r.Form.Get("fname"),
			r.Form.Get("lname"),
			r.Form.Get("mname"),
			"",
			r.Form.Get("department"),
		}

		err := c.Db.AddUser(&user)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
