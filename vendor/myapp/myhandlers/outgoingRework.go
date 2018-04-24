package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
	"strconv"
	"time"
)

// OutgoingRework ...
func (c *AppContext) OutgoingRework(w http.ResponseWriter, r *http.Request) {
	log.Println("OutgoingRework HandleFunc executing ...")

	if r.Method == "POST" {
		r.ParseForm()

		id, _ := strconv.Atoi(r.Form.Get("id"))

		err := c.Db.SetRework(models.SetRework{id, time.Now(), r.Form.Get("calldown")})
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/outgoing", http.StatusSeeOther)
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
