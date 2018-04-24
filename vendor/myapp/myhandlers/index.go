package myhandlers

import (
	"errors"
	"log"
	"net/http"
)

type pageIndex struct {
	*MyCustomClaims
	Picture string
	Error   string
}

// Index выводит сводную информацию по поручениям
func (c *AppContext) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Index Handler executing...")

	if r.Method == "GET" {
		http.Redirect(w, r, "/incoming", http.StatusSeeOther)

		// key := keyContext("Values")
		// claims := r.Context().Value(key).(*MyCustomClaims)
		// picture := "Здесь будет рисунок на ОШС гимназии"

		// c.Tmpl.ExecuteTemplate(w, "index", pageIndex{claims, picture, ""})
		return

	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
