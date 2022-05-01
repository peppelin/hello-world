package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func infoCurl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world\n"))
	w.Write([]byte("\033[46m\033[30mURI\033[0m " + r.URL.Path + "\n"))
	w.Write([]byte(hostnameCurl()))
	w.Write([]byte(envvarsCurl()))
}

// Returns hostname in colored mode for bash, or an error if the hostname could not be retrieved.
func hostnameCurl() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting the hostname: %v.\n", err)
		return ("Error getting the hostname.\n")
	}
	return ("\033[46m\033[30mHOSTNAME\033[0m " + hostname + "\n")
}

//Loops though all the envvars with APP_ prefix and returns its value colored for bash
func envvarsCurl() string {
	var result = "----------\nENV VARS STARTING BY APP_\n----------\n"
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "APP_") {
			fmt.Println(env)
			envSplit := strings.Split(env, "=")
			result += "\033[46m\033[30m" + envSplit[0] + "\033[0m " + envSplit[1] + "\n"
		}
	}
	return (result)
}
