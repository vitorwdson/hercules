package main

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

func main() {
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

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":3000", nil)
}

func GetTemplate(s string) {
	panic("unimplemented")
}
