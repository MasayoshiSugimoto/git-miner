package logmanager

import (
	"log"
	"time"
)

type Commit struct {
	Project   string
	Commit    string
	Author    string
	Timestamp int
	Timezone  string
	Log       string
}

type LogManager struct {
	commits []*Commit
}

type LogFilter func(commits []*Commit) []*Commit

func (logManager *LogManager) AddCommit(commit *Commit) {
	logManager.commits = append(logManager.commits, commit)
}

func (logManager *LogManager) NbCommitPerDayOfWeek(logFilter LogFilter) [7]int {
	nbCommitsPerDayOfWeek := [7]int{}
	for _, commit := range logFilter(logManager.commits) {
		t := time.Unix(int64(commit.Timestamp), 0)
		index := 0
		switch t.Weekday() {
		case time.Monday:
			index = 0
		case time.Tuesday:
			index = 1
		case time.Wednesday:
			index = 2
		case time.Thursday:
			index = 3
		case time.Friday:
			index = 4
		case time.Saturday:
			index = 5
		case time.Sunday:
			index = 6
		}
		nbCommitsPerDayOfWeek[index]++
	}
	return nbCommitsPerDayOfWeek
}

func (logManager *LogManager) Dump() {
	log.Println("Dumping logManager...")
	for _, commit := range logManager.commits {
		log.Println(commit)
	}
}

func (logManager *LogManager) Repos() []string {
	repoMap := map[string]bool{}
	for _, repo := range logManager.commits {
		repoMap[repo.Project] = true
	}
	repos := []string{}
	for key := range repoMap {
		repos = append(repos, key)
	}
	return repos
}

func RepoFilter(repo string) LogFilter {
	if repo == "" {
		return func(commits []*Commit) []*Commit {
			return commits
		}
	}
	filter := func(commits []*Commit) []*Commit {
		filteredCommits := []*Commit{}
		for _, commit := range commits {
			if commit.Project == repo {
				filteredCommits = append(filteredCommits, commit)
			}
		}
		return filteredCommits
	}
	return filter
}
