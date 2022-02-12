package main1

import (
	"fmt"
	"testing"
)

func TestDeckCount(t *testing.T) {
	got := 81
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

func TestGetNumber(t *testing.T) {
	num := getNumber()
	fmt.Println("num is: ", num)
	if num != 100 {
		t.Errorf("num = %d; want 100", num)
	}
}

/*
RUN CMD: go test

--- FAIL: TestDeckCount (0.00s)
    deck_test.go:11: Abs(-1) = 81; want 1
num is:  10
--- FAIL: TestGetNumber (0.00s)
    deck_test.go:19: num = 10; want 100
FAIL
exit status 1
FAIL    go-udemy-basics/deck    0.504s
*/
