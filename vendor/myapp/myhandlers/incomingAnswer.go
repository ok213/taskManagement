package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
	"strconv"
	"time"
)

// IncomingAnswer ...
func (c *AppContext) IncomingAnswer(w http.ResponseWriter, r *http.Request) {
	log.Println("IncomingAnswer HandleFunc executing ...")

	if r.Method == "POST" {
		r.ParseForm()
		// log.Printf("-----> ID: %v \n", r.Form.Get("id"))
		// log.Printf("-----> Answer: %v \n", r.Form.Get("answer"))
		t := time.Now()
		id, _ := strconv.Atoi(r.Form.Get("id"))
		answ := models.SetAnswer{id, t, r.Form.Get("answer")}

		err := c.Db.SetAnswer(answ)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/incoming", http.StatusSeeOther)
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
