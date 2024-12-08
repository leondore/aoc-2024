package day5

import (
	"reflect"
	"strings"
	"testing"
)

func TestDay5(t *testing.T) {
	ordered, unordered, err := Day5("./test_instructions.txt", "./test_updates.txt")
	wantOrd := 143
	wantUnord := 123

	if err != nil {
		t.Fatal(err)
	}

	if ordered != wantOrd {
		t.Errorf("ordered: got %d , want %d", ordered, wantOrd)
	}

	if unordered != wantUnord {
		t.Errorf("unordered: got %d , want %d", unordered, wantUnord)
	}
}

func TestIsUpdateSorted(t *testing.T) {
	cases := []struct {
		updates  string
		expected bool
	}{
		{"75,47,61,53,29", true},
		{"75,29,13", true},
		{"97,13,75,29,47", false},
	}

	instructions, _ := ParseInstructions("./test_instructions.txt")

	for _, test := range cases {
		t.Run(test.updates, func(t *testing.T) {
			updates := strings.Split(test.updates, ",")
			got := IsUpdateSorted(instructions, updates)

			if got != test.expected {
				t.Errorf("got %v, want %v", got, test.expected)
			}
		})
	}
}

func TestParseInstructions(t *testing.T) {
	cases := []struct {
		path     string
		expected map[string]map[string]int
		name     string
	}{
		{"./test_instructions_1.txt", map[string]map[string]int{
			"47": {"53": 1, "97": -1},
			"53": {"47": -1},
			"97": {"13": 1, "61": 1, "47": 1},
			"13": {"97": -1},
			"61": {"97": -1},
			"75": {"29": 1},
			"29": {"75": -1},
		}, "47|53, 97|13, 97|61, 97|47, 75|29"},
		{"./test_instructions_2.txt", map[string]map[string]int{
			"97": {"29": 1, "53": 1},
			"29": {"97": -1, "53": -1, "61": -1},
			"53": {"29": 1, "61": -1, "97": -1},
			"61": {"53": 1, "29": 1},
		}, "97|29, 53|29, 61|53, 97|53, 61|29"},
		{"./test_instructions_3.txt", map[string]map[string]int{
			"47": {"61": 1, "29": 1},
			"61": {"47": -1, "75": -1},
			"75": {"61": 1, "13": 1},
			"29": {"47": -1},
			"13": {"75": -1, "53": -1},
			"53": {"13": 1},
		}, "47|61, 75|61, 47|29, 75|13, 53|13"},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got, err := ParseInstructions(test.path)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("got %v, want %v", got, test.expected)
			}
		})
	}
}
