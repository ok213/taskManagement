package myhandlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

// OutgoingAccept ...
func (c *AppContext) OutgoingAccept(w http.ResponseWriter, r *http.Request) {
	log.Println("OutgoingAccept HandleFunc executing ...")

	if r.Method == "POST" {
		r.ParseForm()
		// log.Printf("-----> ID: %v \n", r.Form.Get("id"))
		// log.Printf("-----> Answer: %v \n", r.Form.Get("answer"))
		answ := r.Form.Get("accept")
		if answ != "accept" {
			log.Println("Error status!!!")
			http.Redirect(w, r, "/outgoing", http.StatusSeeOther)
			return
		}

		id, _ := strconv.Atoi(r.Form.Get("id"))

		err := c.Db.SetAccept(id)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/outgoing", http.StatusSeeOther)
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
