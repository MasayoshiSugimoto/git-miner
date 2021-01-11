package controller

import (
	"log"
	"masa/gitminer/logmanager"
	"masa/gitminer/ui"
	"net/http"
)

func Start(logManager *logmanager.LogManager) {
	log.Println("Starting controller")

	http.HandleFunc("/gitminer", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /gitminer")
		err := r.ParseForm()
		if err != nil {
			log.Panicf("Failed to parse form: %+v", r)
		}
		log.Printf("form = %+v", r.Form)

		ui.DashboardPage(w, logManager.NbCommitPerDayOfWeek(logmanager.RepoFilter(r.Form.Get("repo"))), logManager.Repos())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
