package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// renderTemplate using html/tamplate
func RenderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./template/" + html)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template :", err)
	}
}
