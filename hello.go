package main

import (
	"fmt"
	"net/http"
	instana "github.com/instana/go-sensor"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"message\":\"Hello world!\"}")
}
