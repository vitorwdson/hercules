package auth

import (
	"database/sql"
	"net/http"

	"github.com/vitorwdson/hercules/htmx"
	"github.com/vitorwdson/hercules/models/user"
	"github.com/vitorwdson/hercules/templates"
)

type IndexTemplateData struct {
	Username       string
	Name           string
	Nickname       string
	Password       string
	UsernameError  string
	NameError      string
	NicknameError  string
	Password1Error string
	Password2Error string
}

func RegisterIndex(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	tmpl, err := templates.GetTemplate("views/register.html")
	if err != nil {
		return err
	}

	var data IndexTemplateData

	if r.Method == "POST" {
		r.ParseForm()

		username := r.FormValue("username")
		name := r.FormValue("name")
		nickname := r.FormValue("nickname")
		password1 := r.FormValue("password1")
		password2 := r.FormValue("password2")

		hasError := false
		data = IndexTemplateData{
			Username: username,
			Name:     name,
			Nickname: nickname,
			Password: password1,
		}

		if username == "" {
			hasError = true
			data.UsernameError = "This field is required"
		}

		if name == "" {
			hasError = true
			data.NameError = "This field is required"
		}

		if nickname == "" {
			hasError = true
			data.NicknameError = "This field is required"
		}

		if password1 == "" {
			hasError = true
			data.Password1Error = "This field is required"
		} else if msg := user.CheckPasswordStrength(password1); msg != "" {
			hasError = true
			data.Password1Error = msg
        }

		if password2 == "" {
			hasError = true
			data.Password2Error = "This field is required"
		} else if password1 != password2 {
			hasError = true
			data.Password2Error = "The passwords must match"
		}

        if !hasError {
            user := user.User{
                Username: username,
                Name: name,
                Nickname: nickname,
            }

            err := user.SetPassword(password1)
            if err != nil {
                return err
            }

            err = user.Save(db)
            if err != nil {
                return err
            }
        }
	}

	return htmx.HxExecuteTemplate(tmpl, w, r, data)
}
