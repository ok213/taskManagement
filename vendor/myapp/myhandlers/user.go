package myhandlers

import (
	"errors"
	// "fmt"
	"log"
	"myapp/models"
	"net/http"
)

type pageUser struct {
	*MyCustomClaims
	User  *models.User
	Error string
}

// User выводит сводную информацию по поручениям
func (c *AppContext) User(w http.ResponseWriter, r *http.Request) {
	log.Println("User Handler executing...")

	if r.Method == "GET" {
		key := keyContext("Values")
		claims := r.Context().Value(key).(*MyCustomClaims)

		values := r.URL.Query()
		if len(values) < 1 {
			log.Println("Url Param 'key' is missing")
			http.Redirect(w, r, "/outgoing", http.StatusSeeOther)
			return
		}

		user, err := c.Db.GetUser(values.Get("email"))
		if err != nil {
			log.Printf("User Handler Error: %v\n", err.Error())
			c.Tmpl.ExecuteTemplate(w, "user", pageUser{claims, nil, err.Error()})
			return
		}

		c.Tmpl.ExecuteTemplate(w, "user", pageUser{claims, user, ""})
		return

	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
