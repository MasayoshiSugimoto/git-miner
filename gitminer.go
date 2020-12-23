package main

import (
	"log"
	"masa/gitminer/controller"
	"masa/gitminer/gitlogparser"
	"net/http"
	"os/exec"
	//"masa/gitminer/ui"
)

func main() {

	go startFileServer()
	// startChrome()

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

func startChrome() {
	cmd := exec.Command("cmd", "/c", "start", "chrome", "http://localhost:8080/gitminer")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Started Chrome on 'localhost:8080'")
}
