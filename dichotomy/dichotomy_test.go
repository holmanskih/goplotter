package dichotomy

import (
	"log"
	"testing"
)

func TestDichotomy(t *testing.T) {
	c, err := Dichotomy(-200, 300)
	if err != nil {
		t.Fatalf("failed to run dichoromy method %s", err)
	}

	log.Print(c)
}
