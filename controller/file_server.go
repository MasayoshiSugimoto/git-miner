package controller

import (
	"log"
	"net/http"
)

func StartFileServer() http.Handler {
	fs := http.FileServer(http.Dir("./static/node_modules"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")

	start := func() {
		err := http.ListenAndServe(":3000", fs)
		if err != nil {
			log.Fatal(err)
		}
	}
	go start()
	return fs
}
