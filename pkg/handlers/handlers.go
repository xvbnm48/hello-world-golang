package handlers

import (
	"net/http"

	"github.com/xvbnm48/hello-world-golang/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home_page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about_page.html")
}
