package htmx

import (
	"html/template"
	"net/http"
)

func IsHxRequest(r *http.Request) bool {
    hx := r.Header.Get("HX-Request")

    return hx == "true"
}

func HxExecuteTemplate(tmpl *template.Template, w http.ResponseWriter, r *http.Request, data any) error {
    templateName := "base"
    if IsHxRequest(r) {
        templateName = "main"
    }

    return tmpl.ExecuteTemplate(w, templateName, data)
}
