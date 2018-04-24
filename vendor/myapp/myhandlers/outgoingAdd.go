package myhandlers

import (
	"errors"
	"fmt"
	"log"
	"myapp/models"
	"net/http"
	"strconv"
	"time"
)

// Executors ...
type Executors struct {
	ID       int
	FullName string
}

type pageOutgoingAdd struct {
	*MyCustomClaims
	DateStart string
	Period    string
	Executors *[]Executors
	Error     string
}

func getUserNamesDept(l string, f string, m string, d string) string {
	return fmt.Sprintf("%s %s %s (%s)", l, f, m, d)
}

func getExecutors(users []models.UserDeptName) *[]Executors {
	execs := make([]Executors, 0)
	for _, user := range users {
		exec := Executors{user.Id, getUserNamesDept(user.LastName, user.FirstName, user.MiddleName, user.Department)}
		execs = append(execs, exec)
	}

	return &execs
}

// OutgoingAdd выводит информацию по входящим поручениям
func (c *AppContext) OutgoingAdd(w http.ResponseWriter, r *http.Request) {
	log.Println("OutgoingAdd HandleFunc executing ...")
	key := keyContext("Values")
	claims := r.Context().Value(key).(*MyCustomClaims)
	execs, err := c.Db.GetUsersExcept(claims.Email)
	exs := getExecutors(execs)

	if r.Method == "GET" {
		ds := time.Now().Local().Format("02.01.2006")
		pd := time.Now().Local().AddDate(0, 0, 1).Format("02.01.2006")

		var serr string
		if err != nil {
			serr = err.Error()
		}

		pageVals := pageOutgoingAdd{claims, ds, pd, exs, serr}
		c.Tmpl.ExecuteTemplate(w, "outgoingAdd", pageVals)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()
		idExecutor, _ := strconv.Atoi(r.Form.Get("executor"))

		ds, _ := time.Parse("02.01.2006", r.Form.Get("dateStart"))
		pd, _ := time.Parse("02.01.2006", r.Form.Get("period"))
		instruction := models.SetInstruction{claims.ID, idExecutor, ds, pd, r.Form.Get("content")}

		err := c.Db.SetInstruction(instruction)
		if err != nil {
			log.Println(err)

			pageVals := pageOutgoingAdd{claims, r.Form.Get("dateStart"), r.Form.Get("period"), exs, err.Error()}
			c.Tmpl.ExecuteTemplate(w, "outgoingAdd", pageVals)
			return
		}

		http.Redirect(w, r, "/outgoing", http.StatusSeeOther)
		return
	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)
}
