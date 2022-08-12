package  main

import (
	"testing"
	"github.com/jamesstockonj1/go-serial/main"
)

func TestMultiply(t, *testing.T) {
	answer := Multiply(5, 5)

	if answer != (5 * 5) {
		t.Errorf("Multiply 5x5=%d: want %d", answer, (5 * 5))
	}
}