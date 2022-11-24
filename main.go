package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	content = "Content-type"
	html    = "text/html; charset=utf-8"
	plain   = "text/plain"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(content, html)
		w.Write([]byte("<h1>Hello world</h1>"))
		w.Write([]byte("<strong>URI</strong>: " + r.URL.Path + "</br>"))
		returnHostnameHTML(w)
		returnEnvvarsHTML(w)

	})
	http.HandleFunc("/curl/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world\n"))
		w.Write([]byte("\033[46m\033[30mURI\033[0m " + r.URL.Path + "\n"))
		w.Write([]byte(returnHostname()))
		w.Write([]byte(returnEnvvars()))

	})

	log.Fatal(s.ListenAndServe())
}

// Returns hostname in colored mode for bash, or an error if the hostname could not be retrieved.
func returnHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting the hostname: %v.\n", err)
		return ("Error getting the hostname.\n")
	}
	return ("\033[46m\033[30mHOSTNAME\033[0m " + hostname + "\n")
}

//Loops though all the envvars with APP_ prefix and returns its value colored for bash
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

// Returns hostname in HTML format, or an error if the hostname could not be retrieved.

func returnHostnameHTML(w http.ResponseWriter) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting the hostname: %v.\n", err)
		w.Write([]byte("Error getting the hostname.</br>"))
	} else {
		w.Write([]byte("<strong>HOSTNAME</strong>: " + hostname + "</br>"))
	}
}

//Loops though all the envvars with APP_ prefix and returns its value in HTML format
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
