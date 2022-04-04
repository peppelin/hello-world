package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

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
		returnURIHTML(w, r)
		returnHostnameHTML(w)
		returnEnvvarsHTML(w)

	})
	r.Get("/curl/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world\n"))
		w.Write([]byte(returnURI(r)))
		w.Write([]byte(returnHostname()))
		w.Write([]byte(returnEnvvars()))

	})

	http.ListenAndServe(":8080", r)
}

func returnURI(r *http.Request) string {
	return ("\033[46m\033[30mURI\033[0m " + r.URL.Path + "\n")
}
func returnHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting the hostname: %v.\n", err)
		return ("Error getting the hostname.\n")
	}
	return ("\033[46m\033[30mHOSTNAME\033[0m " + hostname + "\n")
}
func returnEnvvars() string {
	var result = "\tENV VARS\n"
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "APP_") {
			fmt.Println(env)
			envSplit := strings.Split(env, "=")
			result += "\033[46m\033[30m" + envSplit[0] + "\033[0m " + envSplit[1] + "\n"
		}
	}
	return (result)
}
func returnURIHTML(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<strong>URI</strong>: " + r.URL.Path + "</br>"))
}

func returnHostnameHTML(w http.ResponseWriter) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting the hostname: %v.\n", err)
		w.Write([]byte("Error getting the hostname.</br>"))
	} else {
		w.Write([]byte("<strong>HOSTNAME</strong>: " + hostname + "</br>"))
	}
}

func returnEnvvarsHTML(w http.ResponseWriter) {
	var result = "<h2>ENV VARS</h2>"
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "APP_") {
			fmt.Println(env)
			envSplit := strings.Split(env, "=")
			result += "<strong>" + envSplit[0] + "</strong>: " + envSplit[1] + "</br>"
		}
	}
	w.Write([]byte(result))
}
