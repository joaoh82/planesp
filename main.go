package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Planesp Consultoria</h1>")
	} else if r.URL.Path == "/contato" {
		fmt.Fprint(w, "Para entrar em contato, por favor envie um email para <a href=\"mailto:planesp@planesp.com.br\">planesp@planesp.com.br</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>OPS! Não conseguimos achar a página que você está procurando.</h1>")
	}
}
