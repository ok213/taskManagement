package myhandlers

import (
	"errors"
	"log"
	"myapp/models"
	"net/http"
)

type pageAdmin struct {
	*MyCustomClaims
	Users       *[]models.User
	Roles       *[]models.Role
	Departments *[]models.Department
	Error       string
}

// Admin выводит сводную информацию по поручениям
func (c *AppContext) Admin(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Handler executing...")

	if r.Method == "GET" {
		key := keyContext("Values")
		claims := r.Context().Value(key).(*MyCustomClaims)

		users, err := c.Db.GetUsers()
		if err != nil {
			log.Printf("Admin Handler Error: %v\n", err.Error())
			c.Tmpl.ExecuteTemplate(w, "admin", pageAdmin{claims, nil, nil, nil, err.Error()})
			return
		}

		var roles *[]models.Role
		roles, err = c.Db.GetRoles()
		if err != nil {
			log.Printf("Admin Handler Error: %v\n", err.Error())
			c.Tmpl.ExecuteTemplate(w, "admin", pageAdmin{claims, nil, nil, nil, err.Error()})
			return
		}

		var departments *[]models.Department
		departments, err = c.Db.GetDepartments()
		if err != nil {
			log.Printf("Admin Handler Error: %v\n", err.Error())
			c.Tmpl.ExecuteTemplate(w, "admin", pageAdmin{claims, nil, nil, nil, err.Error()})
			return
		}

		c.Tmpl.ExecuteTemplate(w, "admin", pageAdmin{claims, users, roles, departments, ""})
		return

	}

	http.Error(w, errors.New("Некорректный запрос от клиента").Error(), http.StatusInternalServerError)

}
