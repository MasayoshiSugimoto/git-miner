package ui

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"masa/gitminer/logmanager"
)

type HtmlGen struct {
	pageTemplate *template.Template
}

type Page struct {
	Body string
}

func NewHtmlGen() *HtmlGen {
	pageTemplate, err := template.New("name").Parse(`
		<html>
		<head>
		</head>
	
		<body>
			{{.Body}}
		</body>
		</html>
	`)
	if err != nil {
		panic(err)
	}

	return &HtmlGen{
		pageTemplate: pageTemplate,
	}
}

func (htmlGen *HtmlGen) Page(w io.Writer, page *Page) {
	err := htmlGen.pageTemplate.Execute(w, page)
	if err != nil {
		log.Printf(`Failed to render page: %v with error: %v\n`, page, err)
	}
}

func dashboard(board []float64) string {
	tmp, err := template.New("dashboard").Parse(`
		<style>
			.chart-bar {
				background-color: black;
			}
		</style>
		<div style="display: flex; align-items: flex-end;">
			{{range .}}<div class="chart-bar" style="height={{.}}%"></div>{{end}}
		</div>
	`)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = tmp.Execute(buf, board)
	if err != nil {
		log.Fatalf("Failed to generate dashboard: %v", err)
		return ""
	}

	return buf.String()
}

func DashboardPageOld(w io.Writer, logManager *logmanager.LogManager) {
	tmp, err := template.New("dashboard").Parse(`
		<html>
		<head>
		</head>
		<body>
			<style>
				.chart-bar {
					background-color: LightBlue;
					margin: 10px;
				}

				.chart {
					width: 100%;
					height: 100%;
					display: flex;
					align-items: flex-end;
					flex-direction: row;
					border: solid;
					border-color: LightBlue;
				}

				.chart-column {
					width: 14.2%;
					display: flex;
					flex-direction: column;
					justify-content: flex-end;
					height: 100%;
				}

				.chart-legend {
					text-align: center;
				}
			</style>
			<div class="chart">
				{{range .}}
					<div class="chart-column">
						<div class="chart-bar" style="height: {{.NbCommitPourcentage}}%;"></div>
						<div class="chart-legend">{{.Day}}</div>
					</div>
				{{end}}
			</div>
		</body>
		</html>
	`)
	if err != nil {
		panic(err)
	}

	type record struct {
		NbCommitPourcentage int
		Day                 string
	}

	err = tmp.Execute(w, []record{
		{NbCommitPourcentage: 10, Day: "Monday"},
		{NbCommitPourcentage: 20, Day: "Tuesday"},
		{NbCommitPourcentage: 30, Day: "Wednesday"},
		{NbCommitPourcentage: 40, Day: "Thursday"},
		{NbCommitPourcentage: 50, Day: "Friday"},
		{NbCommitPourcentage: 60, Day: "Saturday"},
		{NbCommitPourcentage: 70, Day: "Sunday"},
	})
	if err != nil {
		log.Fatalf("Failed to generate dashboard page: %v", err)
	}
}

func DashboardPage(w io.Writer, nbCommitPerDayOfWeek [7]int, repos []string) {
	const repoSelector = `
		{{define "repo_selector"}}
		<form action="/gitminer" method="get">
			<label for="repo">Repo:</label>
			<select name="repo" id="repo">
				{{range .}}<option value="{{.Repo}}">{{.RepoName}}</option>{{end}}
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
	}

	type record struct {
		NbCommitPerDay [7]int
		Repos          []repository
	}

	reps := []repository{}
	for _, repo := range repos {
		reps = append(reps, repository{
			Repo:     repo,
			RepoName: repo,
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
