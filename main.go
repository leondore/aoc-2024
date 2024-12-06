package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leondore/aoc-2024/day1"
	"github.com/leondore/aoc-2024/day2"
	"github.com/leondore/aoc-2024/day3"
	"github.com/leondore/aoc-2024/day4"
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
	}
}
