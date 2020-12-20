package main

import (
	"log"
	"masa/gitminer/controller"
	"masa/gitminer/gitlogparser"
	"net/http"
	// "masa/gitminer/ui"
)

func main() {

	go startFileServer()

	logManager := gitlogparser.MineGitLogs()
	controller.Start(logManager)
}

func startFileServer() {
	fs := http.FileServer(http.Dir("./static/node_modules"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
