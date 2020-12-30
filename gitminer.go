package main

import (
	"log"
	"masa/gitminer/controller"
	"masa/gitminer/gitlogparser"
	"masa/gitminer/logmanager"
	"net/http"
	"os/exec"
)

func main() {
	context := appContext{}
	context.init()
}

func startChrome() {
	cmd := exec.Command("cmd", "/c", "start", "chrome", "http://localhost:8080/gitminer")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Started Chrome on 'localhost:8080'")
}

type appContext struct {
	fileServer *http.Handler
	logManager *logmanager.LogManager
}

func (context *appContext) injectFileServer() http.Handler {
	if context.fileServer == nil {
		fs := controller.StartFileServer()
		context.fileServer = &fs
	}
	return *context.fileServer
}

func (context *appContext) injectLogManager() *logmanager.LogManager {
	if context.logManager == nil {
		context.logManager = gitlogparser.MineGitLogs()
	}
	return context.logManager
}

func (context *appContext) init() {
	context.injectFileServer()
	context.injectLogManager()
	controller.Start(context.logManager)
}
