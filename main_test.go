package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	tokenString := returnSearchWords("^token1 ^token2 ~token3", "^")
	if len(tokenString) != 2 {
		t.Fail()
	}
	fmt.Println(tokenString)

	tokenString = returnSearchWords("^token1 ^token2 ~token3", "~")
	if len(tokenString) != 1 {
		t.Fail()
	}
	fmt.Println(tokenString)
}
