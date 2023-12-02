package main

import (
	"aoc2023/aoc"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Program requires day parameter")
	}
	day := os.Args[1]
	switch day {
	case "1":
		aoc.Day1("data/day1.txt")
	case "2":
		aoc.Day2("data/day2.txt")
	default:
		log.Fatal("Day not implemented")
	}
}
