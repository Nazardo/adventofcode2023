package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Constraints
	// 12 red cubes, 13 green cubes, and 14 blue cubes
	const MaxRed int = 12
	const MaxGreen int = 13
	const MaxBlue int = 14

	var sumGameValids int = 0
	var sumSetPowers int = 0

	// Line example
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		colonIdx := strings.IndexByte(line, ':')
		if colonIdx < 0 {
			continue
		}
		gameId, _ := strconv.Atoi(line[5:colonIdx])
		extractions := strings.Split(line[colonIdx+1:], ";")
		gameValid := true
		var gameRed int = 0
		var gameGreen int = 0
		var gameBlue int = 0
		for i := range extractions {
			fields := strings.FieldsFunc(extractions[i], func(r rune) bool {
				return r == ' ' || r == ','
			})
			// foreach color in an extraction
			for i = 0; i < len(fields)/2; i++ {
				number, _ := strconv.Atoi(fields[i*2])
				color := fields[(i*2)+1]
				switch color {
				case "red":
					if number > MaxRed {
						gameValid = false
					}
					if number > gameRed {
						gameRed = number
					}
				case "green":
					if number > MaxGreen {
						gameValid = false
					}
					if number > gameGreen {
						gameGreen = number
					}
				case "blue":
					if number > MaxBlue {
						gameValid = false
					}
					if number > gameBlue {
						gameBlue = number
					}
				}
			}
		}
		sumSetPowers += gameRed * gameGreen * gameBlue
		if gameValid {
			sumGameValids += gameId
		}
	}
	fmt.Println("Day2.1:", sumGameValids)
	fmt.Println("Day2.2:", sumSetPowers)
}
