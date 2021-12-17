package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./template/" + html)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template :", err)
	}
}
