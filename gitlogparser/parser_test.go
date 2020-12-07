package gitlogparser

import (
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	const rawLogs = `commit c929894b81644c44800f805d352a93b644e2cc90
tree 4618cd599e88d7ac59a82e86f40f2a389d565ade
parent 3c38f5da65dc3d31dc3a7c2594cd49597374987d
author Masayoshi Sugimoto <sugimoto.massayoshi@gmail.com> 1606838152 +0900
committer Masayoshi Sugimoto <sugimoto.massayoshi@gmail.com> 1606838152 +0900

	Adding source files

commit 3c38f5da65dc3d31dc3a7c2594cd49597374987d
tree 11c1afb6afa63fa1a43a9b2225defaa18b274afd
author Masayoshi Sugimoto <sugimoto.massayoshi@gmail.com> 1606838128 +0900
committer Masayoshi Sugimoto <sugimoto.massayoshi@gmail.com> 1606838128 +0900

	First commit: Add module file

`

	parser := LogParser{}

	expecteds := []struct {
		commit    string
		author    string
		timestamp int
		timezone  string
		log       string
	}{
		{
			"c929894b81644c44800f805d352a93b644e2cc90",
			"Masayoshi Sugimoto",
			1606838152,
			"+0900",
			"Adding source files",
		},
		{
			"3c38f5da65dc3d31dc3a7c2594cd49597374987d",
			"Masayoshi Sugimoto",
			1606838128,
			"+0900",
			"First commit: Add module file",
		},
	}

	for _, line := range strings.Split(rawLogs, "\n") {
		parser.readLine([]byte(line))
	}

	t.Logf(`parser = %+v`, parser)

	if len(parser.commits) != 2 {
		t.Errorf("Got length: %v, instead of: %v", len(parser.commits), 2)
	}

	for i, expected := range expecteds {
		current := parser.commits[i]
		if current.commit != expected.commit {
			t.Errorf("Got commit: %v, instead of: %v", current.commit, expected.commit)
		}
		if current.author != expected.author {
			t.Errorf("Got author: %v, instead of: %v", current.author, expected.commit)
		}
		if current.timestamp != expected.timestamp {
			t.Errorf("Got timestamp: %v, instead of: %v", current.timestamp, expected.timestamp)
		}
	}

	if parser.current != 1 {
		t.Errorf("Got current: %v, instead of: %v", parser.current, 1)
	}
}

func expect(t *testing.T, s1 string, s2 string) {
	if s1 != s2 {
		t.Errorf("Got: %v, instead of: %v", s1, s2)
	}
}
