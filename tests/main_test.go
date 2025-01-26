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

func TestSub(t *testing.T) {
	result := utils.Sub(3, 4)
	expected := -1
	if result != expected {
		t.Errorf("Sub(3, 4) = %d; want %d", result, expected)
	}
}
