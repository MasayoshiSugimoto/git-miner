package controller

import (
	"log"
	"masa/gitminer/logmanager"
	"masa/gitminer/ui"
	"net/http"
)

// Start setup and starts the http server which handles page requests
func Start(logManager *logmanager.LogManager) {
	log.Println("Starting controller")

	http.HandleFunc("/gitminer", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /gitminer")
		err := r.ParseForm()
		if err != nil {
			log.Panicf("Failed to parse form: %+v", r)
		}
		log.Printf("form = %+v", r.Form)

		repo := r.Form.Get("repo")

		ui.DashboardPage(
			w,
			logManager.NbCommitPerDayOfWeek(logmanager.RepoFilter(repo)),
			logManager.Repos(),
			repo,
		)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
