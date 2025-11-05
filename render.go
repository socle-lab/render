package render

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
)

type Render interface {
	Page(w http.ResponseWriter, r *http.Request, opts PageOptions) error
}

type PageOptions struct {
	Variables     interface{}
	Data          interface{}
	View          string
	ViewModel     interface{}
	ComponentFunc func(*TemplateData, interface{}) templ.Component
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
	User            string
	Error           string
	Flash           string
}

type RenderEngine struct {
	Engine     interface{}
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	Session    *scs.SessionManager
}

func (rd *RenderEngine) defaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Secure = rd.Secure
	td.ServerName = rd.ServerName
	td.CSRFToken = nosurf.Token(r)
	td.Port = rd.Port
	if rd.Session.Exists(r.Context(), "userID") {
		td.IsAuthenticated = true
	}
	if rd.Session.Exists(r.Context(), "user") {
		td.User = rd.Session.GetString(r.Context(), "user")
	}
	td.Error = rd.Session.PopString(r.Context(), "error")
	td.Flash = rd.Session.PopString(r.Context(), "flash")
	return td
}
