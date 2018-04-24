package myhandlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

// OutgoingCancel ...
func (c *AppContext) OutgoingCancel(w http.ResponseWriter, r *http.Request) {
	log.Println("OutgoingRework HandleFunc executing ...")

	if r.Method == "GET" {
		ids := r.URL.Query()
		if len(ids) < 1 {
			log.Println("Url Param 'key' is missing")
			http.Redirect(w, r, "/outgoing", http.StatusSeeOther)
			return
		}
		id, _ := strconv.Atoi(ids.Get("ID"))
		t := time.Now()

		err := c.Db.CancelInstruction(id, t)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/outgoing", http.StatusSeeOther)
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
