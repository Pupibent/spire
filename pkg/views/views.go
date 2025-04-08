package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, current string) *View {
	files, err := filepath.Glob("E://spire/web/templates/layouts/*.html")
	if err != nil {
		panic(err)
	}
	files = append(files, current)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
