package main

import (
	"testing"
)

func TestCountTo100(t *testing.T) {
	for i := 1; i <= 100; i++ {
		if i < 1 || i > 100 {
			t.Errorf("Count is out of range: %d", i)
		}
	}
}
