package homecontroller

import (
	"net/http"
	"text/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	tempt, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	tempt.Execute(w, nil)
}
