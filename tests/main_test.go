package main

import (
	"testing"
	"lab-git-project/src/utils"
)

func TestAdd(t *testing.T) {
	result := utils.Add(3, 4)
	expected := 7
	if result != expected {
		t.Errorf("Add(3, 4) = %d; want %d", result, expected)
	}
}
