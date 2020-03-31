package views

import "html/template"

// NewView is exported
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.html")

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

// View is exported
type View struct {
	Template *template.Template
}
