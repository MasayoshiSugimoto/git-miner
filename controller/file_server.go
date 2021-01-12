package controller

import (
	"fmt"
	"log"
	"net/http"
)

// StartFileServer starts the file server which serves additional javascript dependencies
func StartFileServer(port int) http.Handler {
	log.Println("Starting file server...")

	fs := http.FileServer(http.Dir("./static/node_modules"))
	http.Handle("/", fs)

	start := func() {
		log.Printf("Listening on :%v...\n", port)
		err := http.ListenAndServe(fmt.Sprintf(":%v", port), fs)
		if err != nil {
			log.Fatal(err)
		}
	}
	go start()
	return fs
}
