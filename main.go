package main

import (
	"fmt"
	"net/http"
)

const portnumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("starting appication on port %s", portnumber))
	_ = http.ListenAndServe(portnumber, nil)

}
