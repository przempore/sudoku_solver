package main

import "testing"

func TestPrintSquare(t *testing.T) {
	got := PrintSquare()
	want := `
 _  _  _ | _  _  _ | _  _  _
 _  _  _ | _  _  _ | _  _  _
 _  _  _ | _  _  _ | _  _  _
---------|---------|---------
 _  _  _ | _  _  _ | _  _  _
 _  _  _ | _  _  _ | _  _  _
 _  _  _ | _  _  _ | _  _  _
---------|---------|---------
 _  _  _ | _  _  _ | _  _  _
 _  _  _ | _  _  _ | _  _  _
 _  _  _ | _  _  _ | _  _  _
`
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestFileRead(t *testing.T) {
	got := Read_config("test_config_file.txt")
	want := `___3__4__
9___1____
4_8_7_5_3
___6__9__
___157___
__5__2___
2_4___6_5
____2___8
__67_8___
`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
