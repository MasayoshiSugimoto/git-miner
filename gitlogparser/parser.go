package gitlogparser

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

func MineGitLogs() *LogParser {
	parser := newLogParser()
	consumeLogs(parser)
	return parser
}

func consumeLogs(parser *LogParser) error {
	cmd := exec.Command("git", "log", "--pretty=raw")
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
		parser.readLine(scanner.Bytes())
	}
	fmt.Println("Finished to scan stdout")

	if err := scanner.Err(); err != nil {
		log.Fatal("Failed scan git logs: ", err)
	}

	fmt.Println("Git logs parsed")
	return nil
}

type Commit struct {
	commit    string
	author    string
	timestamp int
	timezone  string
	log       string
}

type LogParser struct {
	commits []*Commit
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
		c := &(Commit{})
		c.commit = commitHash[1]
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
		parser.currentCommit().author = authorMatch[1]
		var err error
		parser.currentCommit().timestamp, err = strconv.Atoi(authorMatch[3])
		if err != nil {
			fmt.Printf("Failed to convert timestamp: %v", authorMatch[3])
		}
		parser.currentCommit().timezone = authorMatch[4]
	}

	commitLogMatch := commitLogPattern.FindStringSubmatch(l)
	if commitLogMatch != nil {
		parser.currentCommit().log = commitLogMatch[1]
	}
}

func (parser *LogParser) currentCommit() *Commit {
	return parser.commits[parser.current]
}
