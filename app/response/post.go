package response

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/gobel-admin-client-example/app/cookie"
	"github.com/bmf-san/gobel-admin-client-example/app/model"
)

// PostIndex is a data for index view.
type PostIndex struct {
	Posts *model.Posts
}

// PostCreate is a data for create view.
type PostCreate struct {
	Flash      *cookie.Flash
	Old        map[string]string
	Categories *model.Categories
	Tags       *model.Tags
	Status     []string
}

// ExecutePostIndex responses a index view.
func (r *Response) ExecutePostIndex(w http.ResponseWriter, p *PostIndex) error {
	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/post/index.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", p); err != nil {
		return err
	}
	return nil
}

// ExecutePostCreate responses a create view.
func (r *Response) ExecutePostCreate(w http.ResponseWriter, p *PostCreate) error {
	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/post/create.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", p); err != nil {
		return err
	}
	return nil
}
