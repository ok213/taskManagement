package myhandlers

import (
	"errors"
	"log"
	"net/http"
)

type pageReports struct {
	*MyCustomClaims
	Text  string
	Error string
}

// Reports выводит сводную информацию по поручениям
func (c *AppContext) Reports(w http.ResponseWriter, r *http.Request) {
	log.Println("Reports Handler executing...")

	if r.Method == "GET" {
		key := keyContext("Values")
		claims := r.Context().Value(key).(*MyCustomClaims)
		text := "Страница отчетов недостроена! Планируется вывести отчеты по поручениям."

		c.Tmpl.ExecuteTemplate(w, "reports", pageReports{claims, text, ""})
		return

	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
