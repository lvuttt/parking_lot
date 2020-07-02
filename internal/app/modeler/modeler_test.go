package modeler

import (
	"testing"
)

func TestReadLines(t *testing.T) {
	file := File{
		FilePath: "testdata/test.txt",
	}
	lines, err := file.readLines()
	if err != nil {
		t.Error("cannot read file")
	}
	if len(lines) != 2 {
		t.Errorf("Number of lines should be 2 but got %v\n", len(lines))
	}
	if lines[0] != "Hello 12" {
		t.Errorf("line should be Hello but got %s\n", lines[0])
	}
	if lines[1] != "World 34" {
		t.Errorf("line should be World but got %s\n", lines[1])
	}
}