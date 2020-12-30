package gitlogparser

import (
	"bufio"
	"fmt"
	"log"
	"masa/gitminer/logmanager"
	"os/exec"
	"regexp"
	"strconv"
)

func MineGitLogs(workingDir string) *logmanager.LogManager {
	parser := newLogParser()
	consumeLogs(parser, workingDir)
	logManager := &logmanager.LogManager{}
	for _, commit := range parser.commits {
		logManager.AddCommit(commit)
	}
	return logManager
}

func consumeLogs(parser *LogParser, workingDir string) error {
	cmd := exec.Command("git", "log", "--pretty=raw")
	cmd.Dir = workingDir
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Failed execute `git log`: ", err)
	}
	fmt.Println("`git log` executed")

	err = cmd.Start()
	if err != nil {
		log.Fatal("Failed to start `git log` command: ", err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		if parser.currentCommit() != nil {
			parser.currentCommit().Project = workingDir
		}
		parser.readLine(scanner.Bytes())
	}
	fmt.Println("Finished to scan stdout")

	if err := scanner.Err(); err != nil {
		log.Fatal("Failed scan git logs: ", err)
	}

	fmt.Println("Git logs parsed")
	return nil
}

type LogParser struct {
	commits []*logmanager.Commit
	current int
}

func newLogParser() *LogParser {
	return &LogParser{current: -1}
}

var commitPattern = regexp.MustCompile(`^commit (.*)$`)
var treePattern = regexp.MustCompile(`^tree (.*)$`)
var parentPattern = regexp.MustCompile(`^parent (.*)$`)
var authorPattern = regexp.MustCompile(`^author (.+) <(.+)> ([0-9]+) (.*)$`)
var commitLogPattern = regexp.MustCompile(`^\t(.*)`)

func (parser *LogParser) readLine(line []byte) {
	l := string(line)
	fmt.Println(l)
	commitHash := commitPattern.FindStringSubmatch(string(l))
	if commitHash != nil {
		c := &(logmanager.Commit{})
		c.Commit = commitHash[1]
		parser.current = len(parser.commits)
		parser.commits = append(parser.commits, c)
		return
	}

	tree := treePattern.FindStringSubmatch(l)
	if tree != nil {
		return // We don't care about trees for now
	}

	parent := parentPattern.FindStringSubmatch(l)
	if parent != nil {
		return // We don't care about parent for now
	}

	authorMatch := authorPattern.FindStringSubmatch(l)
	if authorMatch != nil {
		parser.currentCommit().Author = authorMatch[1]
		var err error
		parser.currentCommit().Timestamp, err = strconv.Atoi(authorMatch[3])
		if err != nil {
			fmt.Printf("Failed to convert timestamp: %v", authorMatch[3])
		}
		parser.currentCommit().Timezone = authorMatch[4]
	}

	commitLogMatch := commitLogPattern.FindStringSubmatch(l)
	if commitLogMatch != nil {
		if len(parser.currentCommit().Log) == 0 {
			parser.currentCommit().Log = commitLogMatch[1]
		} else {
			parser.currentCommit().Log += "\n" + commitLogMatch[1]
		}
	}
}

func (parser *LogParser) currentCommit() *logmanager.Commit {
	if len(parser.commits) == 0 {
		return nil
	}
	return parser.commits[parser.current]
}
