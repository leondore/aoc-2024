package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leondore/aoc-2024/day1"
	"github.com/leondore/aoc-2024/day2"
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
	}
}
