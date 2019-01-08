package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", nil)
}
