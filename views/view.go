package views

import "html/template"

// NewView : function recieves files parse it and return view
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	files = append(files, "views/layouts/header.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

// View : struct which mantains
type View struct {
	Template *template.Template
}
