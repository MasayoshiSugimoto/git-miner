package logmanager

import (
	"log"
	"time"
)

// Commit is a structure containing a git commit.
type Commit struct {
	Project   string
	Commit    string
	Author    string
	Timestamp int
	Timezone  string
	Log       string
}

// LogManager allows to filter and group commits.
type LogManager struct {
	commits []*Commit
}

// LogFilter defines a function which filters list of commits.
type LogFilter func(commits []*Commit) []*Commit

// AddCommit adds a commit to the `LogManager`.
func (logManager *LogManager) AddCommit(commit *Commit) {
	logManager.commits = append(logManager.commits, commit)
}

/*
NbCommitPerDayOfWeek returns the number of commits per day of the week based
on the filter passed as parameter.
*/
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

// Dump prints all commits of `LogManager`.
func (logManager *LogManager) Dump() {
	log.Println("Dumping logManager...")
	for _, commit := range logManager.commits {
		log.Println(commit)
	}
}

// Repos returns the list of repos found from the logs.
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

// RepoFilter filters the logs of a `LogManager`.
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
