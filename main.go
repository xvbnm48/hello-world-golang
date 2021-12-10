package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is about page")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil)

}
