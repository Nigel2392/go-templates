package response

import (
	"html/template"
	"net/http"

	"github.com/Nigel2392/go-templates"
)

var TEMPLATE_MANAGER *templates.Manager

// Template configuration must be set before calling this function!
//
// See the templates package for more information.
func Render(w http.ResponseWriter, templateName string, data any) error {
	if TEMPLATE_MANAGER == nil {
		panic("Template manager is nil, please set the template manager before calling Render()")
	}
	var t, name, err = TEMPLATE_MANAGER.Get(templateName)
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, name, data)
}

// Render a string as a template
func String(w http.ResponseWriter, templateString string, data any) error {
	var t = template.New("string")
	t, err := t.Parse(templateString)
	if err != nil {
		return err
	}
	// Render template
	return t.ExecuteTemplate(w, "string", data)
}
