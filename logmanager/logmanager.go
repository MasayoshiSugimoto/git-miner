package logmanager

import (
	"time"
)

type Commit struct {
	Commit    string
	Author    string
	Timestamp int
	Timezone  string
	Log       string
}

type LogManager struct {
	commits []*Commit
}

func (logManager *LogManager) AddCommit(commit *Commit) {
	logManager.commits = append(logManager.commits, commit)
}

func (logManager *LogManager) NbCommitPerDayOfWeek() [7]int {
	nbCommitsPerDayOfWeek := [7]int{}
	for _, commit := range logManager.commits {
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
