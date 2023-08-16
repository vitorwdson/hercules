package web

import (
	"log"
	"net/http"

	"github.com/vitorwdson/hercules/templates"
)

type Test struct {
    Name string
    Bar int
}

type ViewData struct {
    Name string
    TestList []Test
}

func SetupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := templates.GetTemplate("views/index.html")
        if err != nil {
            log.Print(err.Error())
            http.Error(w, "Error loading template", 500)
        }

        err = tmpl.ExecuteTemplate(w, "base", ViewData {
            Name: "FooBar",
            TestList: []Test {
                {
                    Name: "Test",
                    Bar: 5,
                },
                {
                    Name: "Bazz",
                    Bar: 6,
                },
                {
                    Name: "Fizzz",
                    Bar: 7,
                },
                {
                    Name: "FoooBar",
                    Bar: 8,
                },
            },
        })
        if err != nil {
            log.Print(err.Error())
            http.Error(w, "Error executing template", 500)
        }
	})

}
