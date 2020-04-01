package views

import "html/template"

// NewView is exported
func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/layout.gohtml",
		"views/layouts/footer.gohtml")

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
