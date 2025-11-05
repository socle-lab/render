package render

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

type JetRender struct {
	RenderEngine
}

// JetPage renders a template using the Jet templating engine
func (rd *JetRender) Page(w http.ResponseWriter, r *http.Request, opts PageOptions) error {
	var vars jet.VarMap
	variables := opts.Variables
	data := opts.Data
	jetViews := rd.Engine.(*jet.Set)

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	td = rd.defaultData(td, r)

	t, err := jetViews.GetTemplate(fmt.Sprintf("%s.jet", opts.View))
	if err != nil {
		log.Println(err)
		return err
	}

	if err = t.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
