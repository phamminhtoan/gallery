package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir is exported
	LayoutDir string = "views/layouts/"
	// TemplateExt is exported
	TemplateExt string = ".gohtml"
)

// NewView is exported
func NewView(layout string, files ...string) *View {
	files = append(files,
		layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View is exported
type View struct {
	Template *template.Template
	Layout   string
}

// layoutfiles return a slice of string representing
// the layout files used in our application
func layoutFiles() []string {
	file, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return file
}

// Render is used to render the view with the predefined layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
