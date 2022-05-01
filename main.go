package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/peppelin/hello-world/controllers"
)

func main() {
	s := &http.Server{
		Addr:           "127.0.0.1:8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", controllers.ServeInfo)

	fmt.Printf("Running server at http://%s\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
