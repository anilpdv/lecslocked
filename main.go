package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>This is root page </h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>This is contanct page</h1>")
	} else {
		fmt.Fprint(w, "<h1>page not found</h1> ", "url path is not found ", r.URL.Path)
	}
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", nil)
}
