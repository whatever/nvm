package nvm

import (
	"log"
	"testing"
)

// TestRandomLine tests whether blank
func TestRandomLineReturnsValues(t *testing.T) {
	_, _, err := RandomLine("./bin/example.txt")
	if err != nil {
		t.Fail()
	}
}

// TestRandomLineFailsMissing tests whether missing files are errors
func TestRandomLineFailsMissing(t *testing.T) {
	_, _, err := RandomLine("./whatever.zap")
	if err == nil {
		t.Fail()
	}
}
