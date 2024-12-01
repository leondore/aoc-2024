package utils

import (
	"reflect"
	"testing"
)

const inputPath = "../inputs/test.txt"

func TestProcessInput(t *testing.T) {
	got, err := ProcessInput(inputPath)
	want := []string{
		"3 4",
		"4 3",
		"2 5",
		"1 3",
		"3 9",
		"3 3",
	}

	if err != nil {
		t.Fatal("got unexpected error", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
