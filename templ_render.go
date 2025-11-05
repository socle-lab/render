package render

import (
	"errors"
	"net/http"
)

type TemplRender struct {
	RenderEngine
}

// func (rd *TemplRender) Page(w http.ResponseWriter, r *http.Request, opts PageOptions) error {
// 	data := opts.Data
// 	component, ok := data.(templ.Component)
// 	if !ok {
// 		return errors.New("templ renderer requires data to be of type templ.Component")
// 	}

// 	return component.Render(r.Context(), w)
// }

func (rd *TemplRender) Page(w http.ResponseWriter, r *http.Request, opts PageOptions) error {

	td := &TemplateData{}
	if opts.Data != nil {
		td = opts.Data.(*TemplateData)
	}

	td = rd.defaultData(td, r)

	if opts.ComponentFunc == nil {
		return errors.New("component is not defined")
	}
	f := opts.ComponentFunc

	component := f(td, opts.ViewModel)

	return component.Render(r.Context(), w)
}
