package templates

import (
	// "fmt"
	"os"
	// "errors"
	"html/template"
)

// Tmpl - структура шаблонов
type Tmpl struct {
	*template.Template
}

// LoadTemplate загружает шаблоны
func LoadTemplate() *Tmpl {
	var pathTemplates = "templates" + string(os.PathSeparator) + "*"
	var templates = template.Must(template.ParseGlob(pathTemplates))
	// if len(templates) == 0 {
	// 	return nil, errors.New("No templates!")
	// }
	return &Tmpl{templates}
}
