package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
)

type pageAdminDepartment struct {
	*MyCustomClaims
	Departments *[]models.Department
	Error       string
}

// AdminDepartments ...
func (c *AppContext) AdminDepartments(w http.ResponseWriter, r *http.Request) {
	log.Println("AdminDepartments Handler executing...")

	if r.Method == "GET" {
		key := keyContext("Values")
		claims := r.Context().Value(key).(*MyCustomClaims)

		departments, err := c.Db.GetDepartments()
		if err != nil {
			log.Printf("Admin Handler Error: %v\n", err.Error())
			c.Tmpl.ExecuteTemplate(w, "adminDepartments", pageAdminDepartment{claims, nil, err.Error()})
			return
		}

		c.Tmpl.ExecuteTemplate(w, "adminDepartments", pageAdminDepartment{claims, departments, ""})
		return

	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
