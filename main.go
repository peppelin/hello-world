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
	startBlue	= "\033[46m\033[30m"
	endBlue = "\033[0m"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ua := r.UserAgent()
		if strings.HasPrefix(ua,"curl") {
			returnCURL(w,r)
		}else{
			returnHTML(w, r)
		}
		

	})
	fmt.Println("Starting server at",s.Addr )
	log.Fatal(s.ListenAndServe())
}

// Returns hostname in colored mode for bash, or an error if the hostname could not be retrieved.
func returnHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting the hostname: %v.\n", err)
		return ("Error getting the hostname.\n")
	}
	return (startBlue + "HOSTNAME" + endBlue + " " + hostname + "\n")
}

// Loops though all the envvars with APP_ prefix and returns its value colored for bash
func returnEnvvars() string {
	var result = "\n" + startBlue +"ENV VARS" + endBlue + "\n"
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "APP_") {
			fmt.Println(env)
			envSplit := strings.Split(env, "=")
			result += startBlue + envSplit[0] + endBlue + envSplit[1] + "\n"
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

// Loops though all the envvars with APP_ prefix and returns its value in HTML format
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

func returnHTML(w http.ResponseWriter, r *http.Request){
	w.Header().Set(content, html)
	w.Write([]byte("<h1>Hello world</h1>"))
	w.Write([]byte("<strong>URI</strong>: " + r.URL.Path + "</br>"))
	returnHostnameHTML(w)
	returnEnvvarsHTML(w)
}

func returnCURL(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world\n"))
	w.Write([]byte(startBlue + "URI" + endBlue + " " + r.URL.Path + "\n"))
	w.Write([]byte(returnHostname()))
	w.Write([]byte(returnEnvvars()))
}