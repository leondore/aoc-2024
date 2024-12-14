package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leondore/aoc-2024/day1"
	"github.com/leondore/aoc-2024/day2"
	"github.com/leondore/aoc-2024/day3"
	"github.com/leondore/aoc-2024/day4"
	"github.com/leondore/aoc-2024/day5"
	"github.com/leondore/aoc-2024/day6"
	"github.com/leondore/aoc-2024/day7"
	"github.com/leondore/aoc-2024/day8"
	"github.com/leondore/aoc-2024/day9"
)

const inputPath = "./inputs"

func main() {
	args := os.Args

	if len(args) <= 1 {
		log.Fatal("you must pass the name of the day you want to run")
	}

	switch args[1] {
	case "day1":
		dist, score, err := day1.Day1(inputPath + "/day1.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Day 1 results: -------")
		fmt.Printf("Total Distance: %d\n", dist)
		fmt.Printf("Similarity Score: %d\n", score)
	case "day2":
		safeReports, err := day2.Day2(inputPath + "/day2.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Day 2 results: -------")
		fmt.Printf("Count of safe reports: %d\n", safeReports)
	case "day3":
		res, err := day3.Day3(inputPath + "/day3.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Day 3 results: -------")
		fmt.Printf("Program result: %d\n", res)
	case "day4":
		res, err := day4.Day4(inputPath + "/day4.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Day 4 results: -------")
		fmt.Printf("Word occurrences: %d\n", res)
	case "day5":
		ordered, unordered, err := day5.Day5(inputPath+"/day5_instructions.txt", inputPath+"/day5_updates.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Day 5 results: -------")
		fmt.Printf("Sum of ordered updates: %d\n", ordered)
		fmt.Printf("Sum of un-ordered updates: %d\n", unordered)
	case "day6":
		visited, stuck, err := day6.Day6(inputPath + "/day6.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Day 6 results: -------")
		fmt.Printf("Visited positions: %d\n", visited)
		fmt.Printf("Times guard got stuck: %d\n", stuck)
	case "day7":
		result, err := day7.Day7(inputPath + "/day7.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Day 7 results: -------")
		fmt.Printf("Total calibration result: %d\n", result)
	case "day8":
		result, err := day8.Day8(inputPath + "/day8.txt")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Day 8 results: -------")
		fmt.Printf("Unique antinodes: %d\n", result)
	case "day9":
		result, _ := day9.Day9(inputPath + "/day9.txt")

		fmt.Println("Day 9 results: -------")
		fmt.Printf("Filesystem checksum: %d\n", result)
	}
}
