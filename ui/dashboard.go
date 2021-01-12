package ui

import (
	"html/template"
	"io"
	"log"
)

// DashboardPage generates the main `Dashboard` page and writes it to the writer passed as parameter.
func DashboardPage(w io.Writer, nbCommitPerDayOfWeek [7]int, repos []string, selectedRepo string) {
	const repoSelector = `
		{{define "repo_selector"}}
		<form action="/gitminer" method="get">
			<label for="repo">Repo:</label>
			<select name="repo" id="repo">
				{{range .}}<option value="{{.Repo}}" {{if .Selected }}selected{{end}}>{{.RepoName}}</option>{{end}}
			</select>
			<br><br>
			<input type="submit" value="Submit">
		</form>
		{{end}}
	`

	const htmlPage = `
		<html>
			<script src="http://localhost:3000/moment/moment.js"></script>
			<script src="http://localhost:3000/chart.js/dist/Chart.js"></script>
		<head>
		</head>
		<body>
			{{template "repo_selector" .Repos}}

			<canvas id="myChart"></canvas>
			<script>
				var ctx = document.getElementById('myChart').getContext('2d');
				var myChart = new Chart(ctx, {
					type: 'bar',
					data: {
						labels: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
						datasets: [{
							label: 'Number Of Commit Per Day Of The Week',
							data: {{.NbCommitPerDay}},
							borderWidth: 1
						}]
					},
					options: {
						scales: {
							yAxes: [{
								ticks: {
									beginAtZero: true
								}
							}]
						}
					}
				});
			</script>
		</body>
		</html>
	`

	tmp := template.Must(template.Must(template.New("dashboard").Parse(htmlPage)).Parse(repoSelector))

	type repository struct {
		Repo     string
		RepoName string
		Selected bool
	}

	type record struct {
		NbCommitPerDay [7]int
		Repos          []repository
	}

	reps := []repository{
		{
			Repo:     "",
			RepoName: "All",
			Selected: selectedRepo == "",
		},
	}
	for _, repo := range repos {
		reps = append(reps, repository{
			Repo:     repo,
			RepoName: repo,
			Selected: repo == selectedRepo,
		})
	}

	err := tmp.Execute(w, record{
		NbCommitPerDay: nbCommitPerDayOfWeek,
		Repos:          reps,
	})
	if err != nil {
		log.Fatalf("Failed to generate dashboard page: %v", err)
	}
}
