package response

import (
	"html/template"
	"net/http"
)

// HomeIndex is a data for index view.
type HomeIndex struct {
	UserID    string
	CSRFToken string
}

// ExectuteHomeIndex responses a index view.
func (r *Response) ExectuteHomeIndex(w http.ResponseWriter, h *HomeIndex) error {
	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/home/index.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", h); err != nil {
		return err
	}
	return nil
}
