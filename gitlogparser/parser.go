package gitlogparser

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func parse(logs string) int {
	return 1
}

func ReadLogs() {
	cmd := exec.Command("git", "log")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(stdout)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	s := ""
	for line, isPrefix, err := reader.ReadLine(); len(line) > 0; {
		if err != nil {
			log.Fatal("Failed to read git logs: ", err)
		}
		if isPrefix {
			s += string(line)
			continue
		}
		s = string(line)
		fmt.Println(s)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Git logs parsed")
}
