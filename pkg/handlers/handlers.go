package handlers

import (
	"net/http"

	"HTML_TEMPLETES/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {

	render.Rendertemp(w, "home.page.html")

}

func About(w http.ResponseWriter, r *http.Request) {

	render.Rendertemp(w, "about.page.html")

}
