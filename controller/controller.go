package controller

import (
	"log"
	"masa/gitminer/logmanager"
	"masa/gitminer/ui"
	"net/http"
)

func Start(logManager *logmanager.LogManager) {

	http.HandleFunc("/gitminer", func(w http.ResponseWriter, r *http.Request) {

		// htmlGen := ui.NewHtmlGen()

		// htmlGen.Page(w, &ui.Page{Body: "<h1>hello</h1>"})

		ui.DashboardPage(w, logManager.NbCommitPerDayOfWeek())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
