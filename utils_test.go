package main

import (
	"testing"
)

func TestSplitFile(t *testing.T) {
	err := SplitFile("small", 3)
	if err != nil {
		t.Error(err)
	}
}
