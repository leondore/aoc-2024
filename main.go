package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leondore/aoc-2024/day1"
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

		fmt.Println("Day 1 resuls: -------")
		fmt.Printf("Total Distance: %d\n", dist)
		fmt.Printf("Similarity Score: %d\n", score)
	}
}
