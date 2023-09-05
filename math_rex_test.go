package main

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestMathRegexep(t *testing.T) {
	contentInput := []byte("10+10=?\n8-6=?\nНесколько слов\n8/2=?")
	contentTest := []byte("10+10=20\n8-6=2\n8/2=4")
	testFile := fstest.MapFS{
		"input.txt":  {Data: contentInput},
		"test.txt":   {Data: contentTest},
		"output.txt": {Data: []byte("")},
	}
	MathRegexp("input.txt", "output.txt")
	want, _ := testFile.ReadFile("test.txt")
	got, _ := testFile.ReadFile("")

	if string(want) != string(got) {
		t.Errorf("MathRegexp work not true, because want %s, got %s", want, got)
	}
}

func Test_readInFile(t *testing.T) {
	fs := fstest.MapFS{
		"input.txt": {Data: []byte("5+9=14\n4+5=9")},
	}

	want := []string{
		"5+9=14",
		"4+5=9",
	}
	got := readInFile("input.txt", fs)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("readLine() - want %v, got %v", want, got)
	}
}
