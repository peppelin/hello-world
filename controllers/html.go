package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	content = "Content-type"
	html    = "text/html; charset=utf-8"
	plain   = "text/plain"
)

func ServeInfo(w http.ResponseWriter, r *http.Request) {
	ua := r.UserAgent()
	if strings.HasPrefix(ua, "curl") {
		infoCurl(w, r)
	} else {
		infoHTML(w, r)
	}
}

func infoHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, html)
	w.Write([]byte("<h1>Hello world</h1>"))
	w.Write([]byte("<strong>URI</strong>: " + r.URL.Path + "</br>"))
	hostnameHTML(w)
	envvarsHTML(w)

}

// Returns hostname in HTML format, or an error if the hostname could not be retrieved.
func hostnameHTML(w http.ResponseWriter) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting the hostname: %v.\n", err)
		w.Write([]byte("Error getting the hostname.</br>"))
	} else {
		w.Write([]byte("<strong>HOSTNAME</strong>: " + hostname + "</br>"))
	}
}

//Loops though all the envvars with APP_ prefix and returns its value in HTML format
func envvarsHTML(w http.ResponseWriter) {
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
