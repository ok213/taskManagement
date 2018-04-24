package myhandlers

import (
	"log"
	"myapp/models"
	"myapp/templates"
)

// AppContext ...
type AppContext struct {
	Db   models.Datastore
	Tmpl *templates.Tmpl
}

// LoadAppContext ...
func LoadAppContext(info models.Info) *AppContext {
	db, err := models.Connect(info)
	if err != nil {
		log.Panic(err)
	}

	var tmpl = templates.LoadTemplate()

	return &AppContext{db, tmpl}
}
