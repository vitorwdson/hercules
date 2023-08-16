package templates

import "html/template"

func GetTemplate(includedPaths ...string) (*template.Template, error) {
    files := []string{
        "./templates/globals/base.html",
    }

    for _, path := range includedPaths {
        files = append(files, "./templates/" + path)
    }

    tmpl, err := template.ParseFiles(files...)
    if err != nil {
        return nil, err
    }

	return tmpl, nil 
}
