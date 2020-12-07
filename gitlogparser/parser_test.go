package gitlogparser

import "testing"

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
	consumeLogs(&parser)

	t.Logf(`parser = %+v`, parser)

	if len(parser.commits) != 2 {
		t.Errorf("Got length: %v, instead of: %v", len(parser.commits), 2)
	}

	if parser.commits[0].commit != "c929894b81644c44800f805d352a93b644e2cc90" {
		t.Errorf("Got commit: %v, instead of: %v", parser.commits[0].commit, "c929894b81644c44800f805d352a93b644e2cc90")
	}
	if parser.commits[0].author != "Masayoshi Sugimoto" {
		t.Errorf("Got author: %v, instead of: %v", parser.commits[0].author, "Masayoshi Sugimoto")
	}
	if parser.commits[0].timestamp != 1606838152 {
		t.Errorf("Got timestamp: %v, instead of: %v", parser.commits[0].timestamp, 1606838152)
	}

	if parser.commits[1].commit != "3c38f5da65dc3d31dc3a7c2594cd49597374987d" {
		t.Errorf("Got commit: %v, instead of: %v", parser.commits[1].commit, "3c38f5da65dc3d31dc3a7c2594cd49597374987d")
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
