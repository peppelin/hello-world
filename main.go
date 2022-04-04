package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	content = "Content-type"
	html    = "text/html; charset=utf-8"
	plain   = "text/plain"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(content, html)

		w.Write([]byte("<h1>Hello world</h1>"))
		w.Write([]byte("<p><strong>URI: </strong>" + r.URL.Path + "</p>"))

		hostname, err := os.Hostname()
		if err != nil {
			fmt.Printf("Error getting the hostname: %v.\n", err)
		}
		w.Write([]byte("<p><strong>HOSTNAME: </strong>" + hostname + "</p>"))
		w.Write([]byte("<p><strong>FLAG ENVVAR: </strong>" + os.Getenv("FLAG") + "</p>"))

	})
	http.ListenAndServe(":8080", r)
}
