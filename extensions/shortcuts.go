package extensions

import (
	"html/template"
	"net/http"
)

// Base extension
// This is the base extension.
type Base struct {
	// The name of the extension.
	// This is used to uniquely identify the extension.
	ExtensionName string

	// The callback that is called when the extension is rendered.
	// This is used to get the template data and the template name.
	// The template name is the name of the template to render.
	// This is not the file name, but the name of the template.
	// IE: {{template "base" .}}
	Callback func(w http.ResponseWriter, r *http.Request) map[string]any
}

//	extension struct.
//
// This is used to render the extension into the base template.
// Avoids having to create a new struct for each extension.
type WithFilename struct {
	Base
	// The file name for the extension.
	// This is the name of the template to render.
	FileName string
}

// Returns the name of the extension.
func (s *WithFilename) Name() string {
	return s.ExtensionName
}

// Returns the file name of the extension.
// This is the name of the template to render.
// The template will be fetched from a template.Manager.
func (s *WithFilename) Filename() string {
	return s.FileName
}

// Returns the template data for the extension.
func (s *WithFilename) View(w http.ResponseWriter, r *http.Request) map[string]any {
	return s.Callback(w, r)
}

type WithTemplate struct {
	Base
	HTMLTemplate *template.Template
}

func (s *WithTemplate) Template(w http.ResponseWriter, r *http.Request) *template.Template {
	return s.HTMLTemplate
}

// Returns the name of the extension.
func (s *WithTemplate) Name() string {
	return s.ExtensionName
}

// Returns the template data for the extension.
func (s *WithTemplate) View(w http.ResponseWriter, r *http.Request) map[string]any {
	return s.Callback(w, r)
}

type WithStrings struct {
	Base
	HTMLString string
}

func (s *WithStrings) String(w http.ResponseWriter, r *http.Request) string {
	return s.HTMLString
}

// Returns the name of the extension.
func (s *WithStrings) Name() string {
	return s.ExtensionName
}

// Returns the template data for the extension.
func (s *WithStrings) View(w http.ResponseWriter, r *http.Request) map[string]any {
	return s.Callback(w, r)
}
