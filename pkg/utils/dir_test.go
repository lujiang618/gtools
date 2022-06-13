package utils

import "testing"

func TestCleanDir(t *testing.T) {
	if err := CleanDir("./test"); err != nil {
		t.Fatal(err)
	}
}
