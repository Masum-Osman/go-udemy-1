package main1

import "testing"

func TestDeckCount(t *testing.T) {
	got := -1
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

/*
RUN CMD: go test
*/
