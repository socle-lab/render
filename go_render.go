package render

import (
	"fmt"
	"net/http"
	"text/template"
)

type GoRender struct {
	RenderEngine
}

func (rd *GoRender) Page(w http.ResponseWriter, r *http.Request, opts PageOptions) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", rd.RootPath, opts.View))
	if err != nil {
		return err
	}
	data := opts.Data
	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &td)
	if err != nil {
		return err
	}

	return nil
}
