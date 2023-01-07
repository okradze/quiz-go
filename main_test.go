package main

import (
	"testing"
)

func TestParseCSVLines(t *testing.T) {
	lines := [][]string{
		{"1+1", "2"},
		{"2+2", "4"},
	}

	problems := parseCSVLines(lines)

	for i, problem := range problems {
		if problem.question != lines[i][0] || problem.answer != lines[i][1] {
			t.Fatalf("Lines %s got parsed into %s", lines, problems)
		}
	}
}
