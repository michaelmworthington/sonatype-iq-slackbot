package main

import (
	"fmt"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
