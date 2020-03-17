package response

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/gobel-admin-client-example/app/cookie"
)

// AuthIndex is a data for index view.
type AuthIndex struct {
	Flash     *cookie.Flash
	CSRFToken string
	Old       map[string]string
}

// ExecuteAuthIndex responses a index view.
func (r *Response) ExecuteAuthIndex(w http.ResponseWriter, i *AuthIndex) error {
	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/auth/login.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", i); err != nil {
		return err
	}
	return nil
}
